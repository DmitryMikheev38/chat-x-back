package auth

import (
	"context"

	"github.com/dm/chat-x-back/models"
)

// UserRepository ...
type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, eMail, password string) (*models.User, error)
}