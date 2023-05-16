package repo

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"github.com/SafetyLink/commons/types"
)

type User interface {
	GetUserByID(ctx context.Context, userID int64) (*authenticationv1.GetUserByIDResponse, error)
	GetUserSecurityByEmail(ctx context.Context, email string) (*types.User, error)
	GetSelf(ctx context.Context, userID int64) (*authenticationv1.GetSelfResponse, error)
}
