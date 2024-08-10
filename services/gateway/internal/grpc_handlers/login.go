package grpchandlers

import (
	"context"

	generated "github.com/kuzin57/grpc-example/services/gateway/internal/generated/proto"
)

func (s *GRPCServer) Login(ctx context.Context, credentials *generated.Credentials) (*generated.Token, error) {
	return s.authService.Login(ctx, credentials)
}
