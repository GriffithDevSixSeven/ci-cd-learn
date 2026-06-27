package postgres

import (
	"ci_cd/internal/domain"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresUserRepo struct {
	db *pgxpool.Pool
}

func Create(db *pgxpool.Pool) *PostgresUserRepo {
	return &PostgresUserRepo{db: db}
}

func (r *PostgresUserRepo) CreateNewUserDB(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (user_name,email,password) VALUES ($1,$2,$3) RETURNING id`
	err := r.db.QueryRow(ctx, query, user.UserName, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return domain.DBErrorUserAlreadyExists
		}
		return fmt.Errorf("postgres error: Ошибка при добавление нового пользователя в базу: %v", err)
	}
	return nil

}

func (r *PostgresUserRepo) CheckCredsUserDB(ctx context.Context, creds *domain.Credentials) error {
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE user_name = $1 AND password = $2)`
	var exists bool
	err := r.db.QueryRow(ctx, query, creds.UserName, creds.Password).Scan(&exists)
	if err != nil {
		return fmt.Errorf("postgres error: Ошибка при проверке кредов: %v", err)
	}
	if !exists {
		return domain.DBErrorInvalidCreds
	}
	return nil
}

func (r *PostgresUserRepo) DeleteUserFromDB(ctx context.Context, user *domain.User) error {
	query := `DELETE FROM users WHERE user_name = $1 AND email = $2 AND password = $3`
	res, err := r.db.Exec(ctx, query, user.UserName, user.Email, user.Password)
	if err != nil {
		return fmt.Errorf("postgres error: Ошибка при удаление пользователя: %v", err)
	}
	if res.RowsAffected() == 0 {
		return domain.DBErrorUserNotFound
	}
	return nil
}
