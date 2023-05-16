package authenticationService

import (
	"context"
	"github.com/SafetyLink/commons/jwt"
	"github.com/SafetyLink/commons/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (as *Authentication) Login(ctx context.Context, email, password string) (string, error) {
	ctx, span := as.tracer.Start(ctx, "authenticationService.Login", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	userResult, err := as.userRepo.GetUserSecurityByEmail(ctx, email)
	if s, ok := status.FromError(err); ok && err != nil {
		if s.Code() == codes.NotFound {
			return "", nil
		} else {
			as.logger.Error("un caught error", zap.Error(err))
			return "", otel.RecordErrorWithAttribute(status.Error(codes.Internal, err.Error()), err, "un caught error", span, attribute.String("user-email", email))
		}
	}

	err = as.bcryptRepo.CompareHashAndPassword(ctx, userResult.Security.Password, password)
	if err != nil {
		return "", status.Error(codes.NotFound, "invalid email or password")
	}

	return jwt.GenerateJwt(userResult.ID), nil

}
