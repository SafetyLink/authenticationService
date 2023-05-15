package controller

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"github.com/SafetyLink/commons/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (as *AuthService) Login(ctx context.Context, in *authenticationv1.LoginRequest) (*authenticationv1.LoginResponse, error) {
	userID, err := as.authenticationSrv.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "%v username or password", err)
	}

	return &authenticationv1.LoginResponse{JwtToken: jwt.GenerateJwt(userID)}, nil
}
