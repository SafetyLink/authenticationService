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
		Security: &types.Security{
			Password:  userSecurity.Password,
			UpdatedAt: userSecurity.AccountSecurityUpdatedAt.Time,
			DeviceID:  userSecurity.AccountDeviceID,
		},
	}
}

func profileToModel(profile []sqlc.GetSelfRow) *types.User {
	var chats []*types.Chat
	if profile[0].ChatID.Int64 != 0 {
		for _, u := range profile {
			chats = append(chats, &types.Chat{
				ChatID:         u.ChatID.Int64,
				UnreadMessages: u.UnreadMessage.Int64,
				LastMessageAt:  u.LastMessageAt.Time,
				Viewed:         u.Viewed.Bool,
				ViewedAt:       u.ViewedAt.Time,
				Users: &types.ChatUser{
					ID:       u.FriendID.Int64,
					Username: u.FriendUsername.String,
					AvatarID: u.FriendAvatarID.Int64,
				},
			})
		}
	} else {
		chats = []*types.Chat{}
	}

	return &types.User{
		Username:  profile[0].Username,
		Email:     profile[0].Email,
		FirstName: profile[0].FirstName.String,
		LastName:  profile[0].LastName.String,
		AvatarID:  profile[0].AvatarID.Int64,
		CreatedAt: profile[0].CreatedAt.Time,
		UpdatedAt: profile[0].UpdatedAt.Time,
		Chat:      chats,
	}

}
