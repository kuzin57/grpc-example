package app

import (
	"fmt"
	"log"
	"net"

	server "github.com/kuzin57/grpc-example/services/gateway/internal/generated/proto"
	grpchandlers "github.com/kuzin57/grpc-example/services/gateway/internal/grpc_handlers"
	"github.com/kuzin57/grpc-example/services/gateway/internal/services/acl"
	"github.com/kuzin57/grpc-example/services/gateway/internal/services/auth"
	"google.golang.org/grpc"
)

type AppServer struct {
	grpcServer *grpc.Server
	port       int
}

func NewServer(port int) *AppServer {
	return &AppServer{
		grpcServer: grpc.NewServer(),
		port:       port,
	}
}

func (a *AppServer) Init() error {
	// TODO: add configs

	var (
		aclService  = acl.NewService()
		authService = auth.NewAuthService()
	)

	grpcServer := grpchandlers.NewServer(aclService, authService)

	server.RegisterGatewayServer(a.grpcServer, grpcServer)

	log.Println("application init successfully!")

	return nil
}

func (a *AppServer) Run() error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return err
	}

	log.Printf("listening on %s\n", l.Addr())

	if err := a.grpcServer.Serve(l); err != nil {
		return err
	}

	return nil
}

func (a *AppServer) Stop() {
	a.grpcServer.GracefulStop()
}
