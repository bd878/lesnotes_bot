package domain

import (
	"context"
	"errors"
)

var ErrExpired = errors.New("token expired")

type ChatGateway interface {
	Signup(ctx context.Context, login, password string) (token string, err error)
	Auth(ctx context.Context, token string) error
	Login(ctx context.Context, login, password string) (token string, err error)
	Delete(ctx context.Context, login, password, token string) error
}