package auth

import (
	"context"

	server "github.com/kuzin57/grpc-example/services/gateway/internal/generated/proto"
)

var (
	_ = AuthService(&service{})
)

type AuthService interface {
	Login(ctx context.Context, credentials *server.Credentials) (*server.Token, error)
	Register(ctx context.Context, personalData *server.UserPersonalData) (*server.Result, error)
}

type service struct {
}

func NewAuthService() AuthService {
	return &service{}
}
