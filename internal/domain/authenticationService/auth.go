package authenticationService

import (
	"context"
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type Repo interface {
	Login(ctx context.Context, email, password string) (int64, error)
	Register(ctx context.Context, username, email, password string) (int64, error)
}

type Authentication struct {
	userRepo   repo.User
	bcryptRepo repo.BcryptRepo
	logger     *zap.Logger
	tracer     trace.Tracer
}

func NewAuthenticationService(userRepo repo.User, bcryptRepo repo.BcryptRepo, tracer trace.Tracer) Repo {
	return &Authentication{
		userRepo:   userRepo,
		bcryptRepo: bcryptRepo,
		tracer:     tracer,
	}
}
