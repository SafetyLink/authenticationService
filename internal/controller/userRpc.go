package controller

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (as *UserService) GetUserByID(ctx context.Context, in *authenticationv1.GetUserByIDRequest) (*authenticationv1.GetUserByIDResponse, error) {
	user, err := as.userRepo.GetUserByID(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}

	return &authenticationv1.GetUserByIDResponse{
		UserId:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		AvatarId:  user.AvatarID,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (as *UserService) GetSelf(ctx context.Context, in *authenticationv1.GetSelfRequest) (*authenticationv1.GetSelfResponse, error) {
	user, err := as.userRepo.GetSelf(ctx, in.GetUserId())
	if err != nil {
		return nil, err
	}
	var chat []*authenticationv1.Chat

	for _, c := range user.Chat {
		chat = append(chat, &authenticationv1.Chat{
			ChatId:        c.ChatID,
			UnreadMessage: c.UnreadMessages,
			LastMessageAt: timestamppb.New(c.LastMessageAt),
			Viewed:        c.Viewed,
			ViewedAt:      timestamppb.New(c.ViewedAt),
			UserId:        c.Users.ID,
			Username:      c.Users.Username,
			AvatarId:      c.Users.AvatarID,
		})
	}

	return &authenticationv1.GetSelfResponse{
		UserId:    user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		AvatarId:  user.AvatarID,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		Chats:     chat,
	}, nil
}
