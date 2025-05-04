package postgres

import (
	"context"

	"github.com/g0shi4ek/store/internal/store/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pg *pgxpool.Pool
}

func NewUserRepository (pgx * pgxpool.Pool) domain.IUserRepository{
	return &UserRepository{pg:pgx}
}

func (r * UserRepository) Create(ctx context.Context, user * domain.User) error{
	query := "INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3) RETURNING id"
	err := r.pg.QueryRow(ctx, query, user.Username, user.PasswordHash, user.Role).Scan(&user.Id)
	return err
}

func (r * UserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error){
	var newUser domain.User
	query := "SELECT id, username, password_hash, role FROM users WHERE username = $1"
	err := r.pg.QueryRow(ctx, query, username).Scan(&newUser.Id, &newUser.Username, &newUser.PasswordHash, &newUser.Role)
	if err != nil{
		return nil, err
	}
	return &newUser, nil
}