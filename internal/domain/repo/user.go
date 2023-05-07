package repo

import (
	"context"
	"github.com/SafetyLink/commons/types"
)

type User interface {
	GetUserByID(ctx context.Context, userID int64) (*types.User, error)
	GetUserSecurityByEmail(ctx context.Context, email string) (*types.User, error)
	GetSelf(ctx context.Context, userID int64) (*types.User, error)
}
