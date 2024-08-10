package auth

import (
	"context"

	"github.com/kuzin57/grpc-example/services/gateway/internal/consts"
	server "github.com/kuzin57/grpc-example/services/gateway/internal/generated/proto"
)

func (s *service) Register(ctx context.Context, personalData *server.UserPersonalData) (*server.Result, error) {
	return &server.Result{
		Code: consts.OperationSuccessCode,
	}, nil
}
