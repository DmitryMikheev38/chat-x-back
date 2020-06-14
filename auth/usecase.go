package auth

import (
	"context"
)

// UseCase ...
type UseCase interface {
	SignUpConfirm(ctx context.Context, token string) error
	SignUp(ctx context.Context, firstName, lastName, eMail, password string) error
	SingnIn(ctx context.Context, eMail, password string) (string, error)
}
