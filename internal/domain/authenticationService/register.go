package authenticationService

import (
	"context"
	"github.com/SafetyLink/commons/errors"
)

func (as *Authentication) Register(ctx context.Context, username, email, password string) (int64, error) {

	//ctx, span := as.tracer.Start(ctx, "authenticationSrv.register")
	//defer span.End()

	userByEmail, err := as.userRepo.GetUserSecurityByEmail(ctx, email)
	if errors.Is(err, errors.ErrNotFound) {
		return -1, errors.ErrNotFound
	}
	if userByEmail != nil {
		return -1, errors.New("email already exist")
	}

	userByUsername, err := as.userRepo.GetUserSecurityByEmail(ctx, username)
	if errors.Is(err, errors.ErrNotFound) {
		return -1, errors.ErrNotFound
	}
	if userByUsername != nil {
		return -1, errors.New("username already exist")
	}

	_, err = as.bcryptRepo.GenerateFromPassword(ctx, password)
	if err != nil {
		return -1, err
	}

	userID := int64(1)

	return userID, nil

}
