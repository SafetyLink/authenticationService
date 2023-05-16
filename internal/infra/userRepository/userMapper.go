package userRepository

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql/sqlc"
	"github.com/SafetyLink/commons/types"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func userByIDToModel(userResult sqlc.GetUserByIDRow) *authenticationv1.GetUserByIDResponse {
	return &authenticationv1.GetUserByIDResponse{
		Username:  userResult.Username,
		Email:     userResult.Email,
		Firstname: userResult.FirstName.String,
		Lastname:  userResult.LastName.String,
		AvatarId:  userResult.AvatarID.Int64,
		CreatedAt: timestamppb.New(userResult.CreatedAt.Time),
		UpdatedAt: timestamppb.New(userResult.UpdatedAt.Time),
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

func selfToModel(selfProfile []sqlc.GetSelfRow) *authenticationv1.GetSelfResponse {
	var chats []*authenticationv1.Chat

	if selfProfile[0].ChatID.Int64 != 0 {
		for _, u := range selfProfile {
			chats = append(chats, &authenticationv1.Chat{
				ChatId:        u.ChatID.Int64,
				UnreadMessage: u.UnreadMessage.Int64,
				LastMessageAt: timestamppb.New(u.LastMessageAt.Time),
				Viewed:        u.Viewed.Bool,
				ViewedAt:      timestamppb.New(u.ViewedAt.Time),
				UserId:        u.FriendID.Int64,
				Username:      u.FriendUsername.String,
				AvatarId:      u.FriendAvatarID.Int64,
			})
		}
	} else {
		chats = []*authenticationv1.Chat{}
	}

	return &authenticationv1.GetSelfResponse{
		UserId:    selfProfile[0].ID,
		Username:  selfProfile[0].Username,
		Email:     selfProfile[0].Email,
		Firstname: selfProfile[0].FirstName.String,
		Lastname:  selfProfile[0].LastName.String,
		AvatarId:  selfProfile[0].AvatarID.Int64,
		CreatedAt: timestamppb.New(selfProfile[0].CreatedAt.Time),
		UpdatedAt: timestamppb.New(selfProfile[0].UpdatedAt.Time),
		Chats:     chats,
	}
}
