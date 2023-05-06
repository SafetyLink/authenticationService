package userRepository

import (
	"github.com/SafetyLink/authenticationService/internal/domain/entities"
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql/sqlc"
)

func userByIDToModel(userResult sqlc.GetUserByIDRow) *entities.User {
	return &entities.User{
		Username:  userResult.Username,
		Email:     userResult.Email,
		FirstName: userResult.FirstName.String,
		LastName:  userResult.LastName.String,
		AvatarID:  userResult.AvatarID.Int64,
		CreatedAt: userResult.CreatedAt.Time,
		UpdatedAt: userResult.UpdatedAt.Time,
	}
}

func userSecurityByEmailToModel(userSecurity sqlc.GetUserSecurityByEmailRow) *entities.User {
	return &entities.User{
		ID:        userSecurity.ID,
		Username:  userSecurity.Username,
		Email:     userSecurity.Email,
		FirstName: userSecurity.FirstName.String,
		LastName:  userSecurity.LastName.String,
		AvatarID:  userSecurity.AvatarID.Int64,
		CreatedAt: userSecurity.CreatedAt.Time,
		UpdatedAt: userSecurity.UpdatedAt.Time,
		Security: entities.Security{
			Password:  userSecurity.Password,
			UpdatedAt: userSecurity.UpdatedAt_2.Time,
			DeviceID:  userSecurity.DeviceID,
		},
	}
}
