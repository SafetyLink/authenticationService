package userRepository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/SafetyLink/commons/errors"
	"go.opentelemetry.io/otel/attribute"

	"github.com/SafetyLink/commons/otel"
	"github.com/SafetyLink/commons/types"
)

func (ur *userRepository) GetUserByID(ctx context.Context, userID int64) (*types.User, error) {
	ctx, span := ur.tracer.Start(ctx, "userRepo.getUserByID")
	defer span.End()

	userResult, err := ur.database.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrNotFound
		} else {
			return nil, otel.RecordErrorWithAttribute(errors.ErrInternal, err, "un caught error", span, ur.logger, attribute.Int("user-id", int(userID)))
		}
	}
	return userByIDToModel(userResult), nil
}

func (ur *userRepository) GetUserSecurityByEmail(ctx context.Context, email string) (*types.User, error) {
	ctx, span := ur.tracer.Start(ctx, "userRepo.getUserSecurityByEmail")
	userResult, err := ur.database.GetUserSecurityByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrNotFound
		} else {

			return nil, otel.RecordErrorWithAttribute(errors.ErrInternal, err, "un caught error", span, ur.logger, attribute.String("user-email", email))
		}
	}

	return userSecurityByEmailToModel(userResult), nil
}

func (ur *userRepository) GetSelf(ctx context.Context, userID int64) (*types.User, error) {
	ctx, span := ur.tracer.Start(ctx, "userRepo.getSelf")
	defer span.End()

	userResult, err := ur.database.GetSelf(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrNotFound
		} else {
			return nil, otel.RecordErrorWithAttribute(errors.ErrInternal, err, "un caught error", span, ur.logger, attribute.Int("user-id", int(userID)))
		}
	}

	fmt.Println(userResult)

	return profileToModel(userResult), nil
}
