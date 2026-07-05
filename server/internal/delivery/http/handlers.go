package http

import (
	"ci_cd/internal/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserCrudService interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, creds *domain.Credentials) error
	DeleteUser(ctx context.Context, user *domain.User) error
}
type UserHandler struct {
	service  UserCrudService
	validate *validator.Validate
}

func NewUserHandler(service UserCrudService, validate *validator.Validate) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validate,
	}
}

func (h *UserHandler) RegisterUserHandler(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json" + err.Error()})
		return
	}

	user := domain.User{
		UserName: request.UserName,
		Email:    request.Email,
		Password: request.Password,
	}
	err := h.service.Register(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"registration": "success"})

}

func (h *UserHandler) LoginUserHandler(c *gin.Context) {
	var request LoginUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error invalid json": err.Error()})
		return
	}
	creds := domain.Credentials{
		UserName: request.UserName,
		Password: request.Password,
	}
	err := h.service.Login(c.Request.Context(), &creds)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"bad creds error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"login": "success"})
}

func (h *UserHandler) DeleteUserHandler(c *gin.Context) {
	var request DeleteUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error invalid json": err.Error()})
		return
	}
	user := domain.User{
		UserName: request.UserName,
		Email:    request.Email,
		Password: request.Password,
	}
	err := h.service.DeleteUser(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"delete user error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"delete user": "success"})
}
