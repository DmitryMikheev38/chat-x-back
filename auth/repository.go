package auth

import "github.com/dm/chat-x-back/models"

// UserRepository ...
type UserRepository interface {
	CreateUser(user *models.User) error
	GetUser(eMail, password string) (*models.User, error)
}