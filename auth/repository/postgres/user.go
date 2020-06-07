package postgres

import (
	"context"

	"github.com/dm/chat-x-back/models"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

// UserRepository ...
type UserRepository struct {
	db *pgx.Conn
}

// NewUserRepository ...
func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser ...
func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := "INSERT INTO users (email, password, firstname, lastname) VALUES ($1, $2, $3, $4) RETURNING id;"
	err := ur.db.QueryRow(ctx, query, &user.EMail, &user.Password, &user.FirstName, &user.LastName).Scan(&user.ID)
	if err != nil {
		return errors.Wrap(err, "CreateUser #1: ")
	}
	return nil
}

func (ur *UserRepository) GetUser(ctx context.Context, eMail, password string) (*models.User, error) {
	return nil, nil
}
