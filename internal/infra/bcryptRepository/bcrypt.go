package bcryptRepository

import (
	"context"
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BcryptRepo struct {
	tracer trace.Tracer
	logger *zap.Logger
}

func NewBcryptRepository(tracer trace.Tracer, logger *zap.Logger) repo.BcryptRepo {
	return &BcryptRepo{
		tracer: tracer,
		logger: logger,
	}
}

func (br *BcryptRepo) CompareHashAndPassword(ctx context.Context, hashedPassword, password string) error {
	ctx, span := br.tracer.Start(ctx, "bcryptRepo.CompareHashAndPassword")
	defer span.End()

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		br.logger.Error("un caught error", zap.Error(err))
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (br *BcryptRepo) GenerateFromPassword(ctx context.Context, password string) (string, error) {
	ctx, span := br.tracer.Start(ctx, "bcryptRepo.GenerateFromPassword")
	defer span.End()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		br.logger.Error("un caught error", zap.Error(err))
		return "", status.Error(codes.Internal, err.Error())
	}

	return string(hashedPassword), nil
}
