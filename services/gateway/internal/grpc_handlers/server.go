package grpchandlers

import (
	server "github.com/kuzin57/grpc-example/services/gateway/internal/generated/proto"
	"github.com/kuzin57/grpc-example/services/gateway/internal/services/acl"
	"github.com/kuzin57/grpc-example/services/gateway/internal/services/auth"
)

type GRPCServer struct {
	server.UnimplementedGatewayServer

	aclService  acl.ACLService
	authService auth.AuthService
}

func NewServer(
	aclService acl.ACLService,
	authService auth.AuthService,
) *GRPCServer {
	return &GRPCServer{
		aclService:  aclService,
		authService: authService,
	}
}
