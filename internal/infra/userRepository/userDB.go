package userRepository

import (
	"context"
	"github.com/SafetyLink/authenticationService/internal/domain/entities"
)

func (ur userRepository) GetUserByID(ctx context.Context, userID int64) (*entities.User, error) {
	userResult, err := ur.database.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return userIDToModel(userResult), nil
}
