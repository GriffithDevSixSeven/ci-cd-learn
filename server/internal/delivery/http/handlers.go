package http

import (
	"ci_cd/internal/domain"
	"ci_cd/internal/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/jackc/pgx/v5/pgxpool"
)



type UserCrudService interface{
	Register(ctx context.Context,user *domain.User) error
	Login(ctx context.Context,creds *domain.Credentials) error
	DeleteUser(ctx context.Context,user *domain.User) error
}
type UserHandler struct {
	service UserCrudService
	validate *validator.Validate
}

func NewUserHandler(service UserCrudService,validate *validator.Validate) *UserHandler {
	return &UserHandler{
		service: service,
		validate: validate,
	}
}



