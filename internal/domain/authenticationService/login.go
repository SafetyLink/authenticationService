package authenticationService

import (
	"context"
	"github.com/SafetyLink/commons/errors"
	"github.com/SafetyLink/commons/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func (as *Authentication) Login(ctx context.Context, email, password string) (int64, error) {
	ctx, span := as.tracer.Start(ctx, "authenticationService.login", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	userResult, err := as.userRepo.GetUserSecurityByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, errors.ErrNotFound) {
			return -1, errors.ErrInvalid
		} else {
			as.logger.Error("un caught error", zap.Error(err))
			return -1, otel.RecordErrorWithAttribute(errors.ErrInternal, err, "un caught error", span, attribute.String("user-email", email))
		}
	}
	err = as.bcryptRepo.CompareHashAndPassword(ctx, userResult.Security.Password, password)
	if err != nil {
		return -1, errors.ErrInvalid
	}

	return userResult.ID, nil

}
