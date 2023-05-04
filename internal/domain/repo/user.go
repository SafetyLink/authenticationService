package repo

import (
	"context"
	"github.com/SafetyLink/authenticationService/internal/domain/entities"
)

type User interface {
	GetUserByID(ctx context.Context, userID int64) (*entities.User, error)
}
