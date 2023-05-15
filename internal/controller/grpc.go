package controller

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	"github.com/SafetyLink/authenticationService/internal/domain/authenticationService"
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
	"google.golang.org/grpc"
)

type AuthService struct {
	authenticationv1grpc.UnimplementedAuthenticationServiceServer
	authenticationSrv authenticationService.Repo
}

type UserService struct {
	authenticationv1grpc.UnimplementedAuthenticationServiceServer
	userRepo repo.User
}

func RegisterAuthenticationService(s *grpc.Server, authenticationSrv authenticationService.Repo) {
	authenticationv1grpc.RegisterAuthenticationServiceServer(s, &AuthService{
		authenticationSrv: authenticationSrv,
	})
}

func RegisterUserService(s *grpc.Server, userRepo repo.User) {
	authenticationv1grpc.RegisterUserServiceServer(s, &UserService{
		userRepo: userRepo,
	})

}
