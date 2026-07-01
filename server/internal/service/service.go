package service

import (
	"ci_cd/internal/domain"
	"ci_cd/internal/repository"
	"context"
	"fmt"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, user *domain.User) error {
	err := s.repo.CreateNewUserDB(ctx, user)
	if err != nil {
		return fmt.Errorf("service error: Ошибка при создании нового пользователя: %v", err)
	}
	return nil
}

func (s *UserService) Login(ctx context.Context, creds *domain.Credentials) error {
	err := s.repo.CheckCredsUserDB(ctx, creds)
	if err != nil {
		return fmt.Errorf("service error: Ошибка при логине пользователя: %v", err)
	}
	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, user *domain.User) error {
	err := s.repo.DeleteUserFromDB(ctx, user)
	if err != nil {
		return fmt.Errorf("service error: Ошибка при удалении пользователя: %v", err)
	}
	return nil
}
