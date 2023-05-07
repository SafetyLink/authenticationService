package userRepository

import (
	"context"
	"github.com/SafetyLink/commons/types"
)

func (ur userRepository) GetUserByID(ctx context.Context, userID int64) (*types.User, error) {
	userResult, err := ur.database.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userByIDToModel(userResult), nil
}

func (ur userRepository) GetUserSecurityByEmail(ctx context.Context, email string) (*types.User, error) {
	userResult, err := ur.database.GetUserSecurityByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return userSecurityByEmailToModel(userResult), nil
}

func (ur userRepository) GetSelf(ctx context.Context, userID int64) (*types.User, error) {
	userResult, err := ur.database.GetSelf(ctx, userID)
	if err != nil {
		return nil, err
	}

	return profileToModel(userResult), nil
}
