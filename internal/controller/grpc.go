package controller

import (
	"buf.build/gen/go/asavor/safetylink/grpc/go/authentication/v1/authenticationv1grpc"
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (as *UserService) GetUserByID(ctx context.Context, in *authenticationv1.GetUserByIDRequest) (*authenticationv1.GetUserByIDResponse, error) {
	user, err := as.userRepo.GetUserByID(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	return &authenticationv1.GetUserByIDResponse{
		UserId:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		AvatarId:  user.AvatarID,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}
