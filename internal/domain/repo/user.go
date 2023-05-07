package repo

import (
	"context"
	"github.com/SafetyLink/commons/types"
)

type User interface {
	GetUserByID(ctx context.Context, userID int64) (*types.User, error)
}
