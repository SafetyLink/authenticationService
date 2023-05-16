package controller

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
)

func (as *AuthService) Login(ctx context.Context, in *authenticationv1.LoginRequest) (*authenticationv1.LoginResponse, error) {
	jwtToken, err := as.authenticationSrv.Login(ctx, in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, err
	}
	return &authenticationv1.LoginResponse{JwtToken: jwtToken}, nil
}

func (as *AuthService) Register(ctx context.Context, in *authenticationv1.RegisterRequest) (*authenticationv1.RegisterResponse, error) {
	jwtToken, err := as.authenticationSrv.Register(ctx, in.GetUsername(), in.GetEmail(), in.GetPassword())
	if err != nil {
		return nil, err
	}
	return &authenticationv1.RegisterResponse{JwtToken: jwtToken}, nil
}
