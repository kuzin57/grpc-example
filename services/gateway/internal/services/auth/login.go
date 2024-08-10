package auth

import (
	"context"

	server "github.com/kuzin57/grpc-example/services/gateway/internal/generated/proto"
)

func (s *service) Login(ctx context.Context, credentials *server.Credentials) (*server.Token, error) {
	return &server.Token{Token: "blabla", Ttl: 24}, nil
}
