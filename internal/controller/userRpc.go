package controller

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
)

func (as *UserService) GetUserByID(ctx context.Context, in *authenticationv1.GetUserByIDRequest) (*authenticationv1.GetUserByIDResponse, error) {
	user, err := as.userRepo.GetUserByID(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (as *UserService) GetSelf(ctx context.Context, in *authenticationv1.GetSelfRequest) (*authenticationv1.GetSelfResponse, error) {
	self, err := as.userRepo.GetSelf(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	return self, nil
}
