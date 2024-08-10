package grpchandlers

import (
	"context"

	generated "github.com/kuzin57/grpc-example/services/gateway/internal/generated/proto"
)

func (s *GRPCServer) Register(ctx context.Context, personalData *generated.UserPersonalData) (*generated.Result, error) {
	return s.authService.Register(ctx, personalData)
}
