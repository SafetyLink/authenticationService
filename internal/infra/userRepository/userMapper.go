package userRepository

import (
	"github.com/SafetyLink/authenticationService/internal/domain/entities"
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql/sqlc"
)

func userIDToModel(userResult sqlc.User) *entities.User {
	return &entities.User{
		ID:        userResult.ID,
		Username:  userResult.Username,
		Email:     userResult.Email,
		FirstName: userResult.FirstName.String,
		LastName:  userResult.LastName.String,
		AvatarID:  userResult.AvatarID.Int64,
		CreatedAt: userResult.CreatedAt.Time,
		UpdatedAt: userResult.UpdatedAt.Time,
	}
}
