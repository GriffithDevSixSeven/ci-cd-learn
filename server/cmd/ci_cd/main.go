package main

import (
	"ci_cd/configs"
	"ci_cd/internal/delivery/http"
	"ci_cd/internal/repository/postgres"
	"ci_cd/internal/service"
	"context"
	"fmt"
	"log"
	netHttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	dbPool, err := pgxpool.New(ctx, configs.GetDBUrl())
	if err != nil {
		log.Fatalf("pool err: Не удалось подключиться к базе: %v", err)
	}
	defer dbPool.Close()
	if err := dbPool.Ping(ctx); err != nil {
		log.Fatalf("pool ping err: База данных не работает:2 %v", err)
	}
	log.Println("pool: Успешное подключение!")
	validate := validator.New()
	userRepo := postgres.Create(dbPool)
	userService := service.NewUserService(userRepo)
	userHanlder := http.NewUserHandler(userService, validate)
	router := gin.Default()

	api := router.Group("/api/v1")
	{
		api.POST("/auth/register", userHanlder.RegisterUserHandler)
		api.POST("/auth/login", userHanlder.LoginUserHandler)
		api.DELETE("/auth/delete_user", userHanlder.DeleteUserHandler)
	}

	srv := &netHttp.Server{
		Addr:    fmt.Sprintf(":%d", configs.GetConfig().Server.Port),
		Handler: router,
	}

	log.Println("Запускаю http!")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != netHttp.ErrServerClosed {
			log.Fatalf("http error: Ошибка при запуске сервера: %v", err)
		}
	}()

	<-ctx.Done()
	log.Println("Получен сигнал остановки, завершаю работу...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Ошибка при плавной остановке сервера: %v", err)
	}

	log.Println("Сервер успешно остановлен. Закрываю пул БД...")

}
