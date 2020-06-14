package usecase

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dm/chat-x-back/auth"
	"github.com/dm/chat-x-back/models"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// AuthUseCase ...
type AuthUseCase struct {
	userRepo       auth.UserRepository
	hashSalt       string
	signInKey      []byte
	expireDuration time.Duration
}

// AuthTokenClaims ...
type AuthTokenClaims struct {
	*jwt.StandardClaims
	models.JWTUserData
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
		IsActive:  false,
	}

	expiresAt := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.New(jwt.SigningMethodHS256)

	token.Claims = &AuthTokenClaims{
		&jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
		models.JWTUserData{
			FirstName: firstName,
			LastName:  lastName,
			EMail:     eMail,
			IsActive:  false,
		},
	}

	tokenStirng, err := token.SignedString([]byte(viper.GetString("appSecret")))
	if err != nil {
		return errors.Wrap(err, "SignUp #1: ")
	}

	user.Token = tokenStirng

	if err := a.userRepo.CreateUser(ctx, user); err != nil {
		return errors.Wrap(err, "SignUp #2: ")
	}

	return nil
}

// SingnIn ...
func (a *AuthUseCase) SingnIn(ctx context.Context, eMail, password string) (string, error) {
	return "", nil
}

// SignUpConfirm ...
func (a *AuthUseCase) SignUpConfirm(ctx context.Context, token string) error {
	return nil
}
