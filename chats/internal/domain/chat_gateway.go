package domain

import (
	"context"
)

type ChatGateway interface {
	Signup(ctx context.Context, login, password string) (token string, err error)
	Login(ctx context.Context, login, password string) (token string, err error)
}