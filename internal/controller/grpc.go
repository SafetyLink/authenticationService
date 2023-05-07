package controller

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
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

type UserService struct {
	authenticationv1grpc.UnimplementedAuthenticationServiceServer
	userRepo repo.User
}

func NewAuthenticationServiceGrpcServer(lis net.Listener, userRepo repo.User) *AuthenticationServiceGrpcServer {
	s := grpc.NewServer()

	authenticationv1grpc.RegisterUserServiceServer(s, &UserService{
		userRepo: userRepo,
	})
	authenticationv1grpc.RegisterAuthenticationServiceServer(s, &AuthService{})

	return &AuthenticationServiceGrpcServer{
		Gs:  s,
		Lis: lis,
	}
}
