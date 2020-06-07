package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dm/chat-x-back/auth"
	"github.com/dm/chat-x-back/models"
	"github.com/pkg/errors"
)

// AuthUseCase ...
type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signInKey      []byte
	expireDuration time.Duration
}

// NewAuthUseCase ...
func NewAuthUseCase(
	userRepo auth.UserRepository,
	hashSalt string,
	signInKey []byte,
	expireDuration time.Duration) *AuthUseCase {
	return &AuthUseCase{
		userRepo:       userRepo,
		hashSalt:       hashSalt,
		signInKey:      signInKey,
		expireDuration: expireDuration,
	}
}

// SignUp ...
func (a *AuthUseCase) SignUp(ctx context.Context, firstName, lastName, eMail, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.hashSalt))

	user := &models.User{
		FirstName: firstName,
		LastName:  lastName,
		EMail:     eMail,
		Password:  fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	if err := a.userRepo.CreateUser(ctx, user); err != nil {
		return errors.Wrap(err, "SignUp #1: ")
	}

	return nil
}

// SingnIn ...
func (a *AuthUseCase) SingnIn(ctx context.Context, eMail, password string) (string, error) {
	return "", nil
}
