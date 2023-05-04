package controller

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	"fmt"
	"github.com/SafetyLink/authenticationService/internal"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type AuthenticationServiceGrpcServer struct {
	Gs  *grpc.Server
	Lis net.Listener
}

type AuthService struct {
	authenticationv1grpc.UnimplementedAuthenticationServiceServer
}

func NewAuthenticationServiceGrpcServer(logger *zap.Logger, config *internal.Config) *AuthenticationServiceGrpcServer {
	lis, err := net.Listen("tcp", config.Grpc.Port)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to start %s grpc server", config.Grpc.ServiceName), zap.Error(err))
	}

	s := grpc.NewServer()

	authenticationv1grpc.RegisterAuthenticationServiceServer(s, &AuthService{})

	return &AuthenticationServiceGrpcServer{
		Gs:  s,
		Lis: lis,
	}

}
