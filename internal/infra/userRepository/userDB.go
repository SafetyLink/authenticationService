package userRepository

import (
	authenticationv1 "buf.build/gen/go/asavor/safetylink/protocolbuffers/go/authentication/v1"
	"context"
	"github.com/SafetyLink/commons/errors"
	"github.com/SafetyLink/commons/otel"
	"github.com/SafetyLink/commons/types"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ur *userRepository) GetUserByID(ctx context.Context, userID int64) (*authenticationv1.GetUserByIDResponse, error) {
	ctx, span := ur.tracer.Start(ctx, "userRepo.postgres.getUserSecurityByEmail", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	userResult, err := ur.database.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			ur.logger.Error("un caught error", zap.Error(err))
			return nil, otel.RecordErrorWithAttribute(status.Error(codes.Internal, err.Error()), err, "un caught error", span, attribute.Int("user-id", int(userID)))
		}
	}
	return userByIDToModel(userResult), nil
}

func (ur *userRepository) GetUserSecurityByEmail(ctx context.Context, email string) (*types.User, error) {
	ctx, span := ur.tracer.Start(ctx, "userRepo.postgres.getUserSecurityByEmail", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	userResult, err := ur.database.GetUserSecurityByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			ur.logger.Error("un caught error", zap.Error(err))
			return nil, otel.RecordErrorWithAttribute(status.Error(codes.Internal, err.Error()), err, "un caught error", span, attribute.String("user-email", email))
		}
	}

	return userSecurityByEmailToModel(userResult), nil
}

func (ur *userRepository) GetSelf(ctx context.Context, userID int64) (*authenticationv1.GetSelfResponse, error) {
	ctx, span := ur.tracer.Start(ctx, "userRepo.postgres.getSelf", trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	userResult, err := ur.database.GetSelf(ctx, userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		} else {
			ur.logger.Error("un caught error", zap.Error(err))
			return nil, otel.RecordErrorWithAttribute(status.Error(codes.Internal, err.Error()), err, "un caught error", span, attribute.Int("user-id", int(userID)))
		}
	}

	return selfToModel(userResult), nil
}
