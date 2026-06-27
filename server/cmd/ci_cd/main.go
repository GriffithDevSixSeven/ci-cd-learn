package main

import (
	"ci_cd/configs"
	"ci_cd/internal/domain"
	"ci_cd/internal/repository/postgres"
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbUrl := configs.GetDBUrl()
	ctx, cannel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cannel()
	pool, err := pgxpool.New(ctx, dbUrl)
	if err != nil {
		log.Fatalf("pool conn: Не удалось подключиться к базе: %v", err)

	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("pool ping: Не удалось сделать пинг: %v", err)
	}
	log.Println("pool conn: Успешное подключение!")
	userRepo := postgres.Create(pool)
	testCreateUser := domain.User{
		UserName: "testName",
		Email:    "testMail",
		Password: "testPassword"}
	log.Println("Попытка добавления данных...")
	err = userRepo.CreateNewUserDB(ctx, &testCreateUser)
	log.Println(err)
	log.Println("Попытка проверки кредов...")

	secondUser := domain.User{
		UserName: "Ivan",
		Email:    "67",
		Password: "123",
	}
	_ = userRepo.CreateNewUserDB(ctx, &secondUser)
	creds := domain.Credentials{
		UserName: "Ivan",
		Password: "123",
	}
	err = userRepo.CheckCredsUserDB(ctx, &creds)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Проверка кредов успешная!")

}
