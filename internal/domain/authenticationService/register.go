package authenticationService

import (
	"context"
	"github.com/SafetyLink/commons/errors"
	"go.opentelemetry.io/otel/trace"
)

func (as *Authentication) Register(ctx context.Context, username, email, password string) (string, error) {
	ctx, span := as.tracer.Start(ctx, "authenticationService.Register", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	userByEmail, err := as.userRepo.GetUserSecurityByEmail(ctx, email)
	if errors.Is(err, errors.ErrNotFound) {
		return "", errors.ErrNotFound
	}
	if userByEmail != nil {
		return "", errors.New("email already exist")
	}

	userByUsername, err := as.userRepo.GetUserSecurityByEmail(ctx, username)
	if errors.Is(err, errors.ErrNotFound) {
		return "", errors.ErrNotFound
	}
	if userByUsername != nil {
		return "", errors.New("username already exist")
	}

	_, err = as.bcryptRepo.GenerateFromPassword(ctx, password)
	if err != nil {
		return "", err
	}

	return "userID", nil

}
