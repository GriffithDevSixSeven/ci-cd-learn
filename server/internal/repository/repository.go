package repository

import (
	"ci_cd/internal/domain"
	"context"
)

type Repository interface {
	CreateNewUserDB(ctx context.Context, user *domain.User) error
	CheckCredsUserDB(ctx context.Context, creds *domain.Credentials) error
	DeleteUserFromDB(ctx context.Context, user *domain.User) error
}
