package userRepository

import (
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql/sqlc"
	"github.com/SafetyLink/commons/types"
)

func userByIDToModel(userResult sqlc.GetUserByIDRow) *types.User {
	return &types.User{
		Username:  userResult.Username,
		Email:     userResult.Email,
		FirstName: userResult.FirstName.String,
		LastName:  userResult.LastName.String,
		AvatarID:  userResult.AvatarID.Int64,
		CreatedAt: userResult.CreatedAt.Time,
		UpdatedAt: userResult.UpdatedAt.Time,
	}
}

func userSecurityByEmailToModel(userSecurity sqlc.GetUserSecurityByEmailRow) *types.User {
	return &types.User{
		ID:        userSecurity.ID,
		Username:  userSecurity.Username,
		Email:     userSecurity.Email,
		FirstName: userSecurity.FirstName.String,
		LastName:  userSecurity.LastName.String,
		AvatarID:  userSecurity.AvatarID.Int64,
		CreatedAt: userSecurity.CreatedAt.Time,
		UpdatedAt: userSecurity.UpdatedAt.Time,
		Security: types.Security{
			Password:  userSecurity.Password,
			UpdatedAt: userSecurity.UpdatedAt_2.Time,
			DeviceID:  userSecurity.DeviceID,
		},
	}
}

func profileToModel(profile []sqlc.GetSelfRow) *types.User {
	var chat []types.Chat

	for _, k := range profile {
		chat = append(chat, types.Chat{
			ChatID:         k.ChatID,
			UnreadMessages: k.UnreadMessage.Int64,
			LastMessageAt:  k.LastMessageAt.Time,
			Viewed:         k.Viewed.Bool,
			ViewedAt:       k.ViewedAt.Time,
			Users: types.User{
				ID:        k.FriendID,
				Username:  k.FriendUsername,
				FirstName: k.FirstName.String,
				AvatarID:  k.FriendAvatarID.Int64,
			},
		})
	}

	return &types.User{
		Username:  profile[0].Username,
		Email:     profile[0].Email,
		FirstName: profile[0].FirstName.String,
		LastName:  profile[0].LastName.String,
		AvatarID:  profile[0].AvatarID.Int64,
		CreatedAt: profile[0].CreatedAt.Time,
		UpdatedAt: profile[0].UpdatedAt.Time,
		Chat:      chat,
	}

}
