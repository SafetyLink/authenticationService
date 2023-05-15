package bcryptRepository

import (
	"context"
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/crypto/bcrypt"
)

type Bcrypt struct {
	tracer trace.Tracer
}

func NewBcryptRepository(tracer trace.Tracer) repo.BcryptRepo {
	return &Bcrypt{
		tracer: tracer,
	}
}

func (b *Bcrypt) CompareHashAndPassword(ctx context.Context, hashedPassword, password string) error {
	ctx, span := b.tracer.Start(ctx, "bcryptRepo.CompareHashAndPassword")
	defer span.End()

	ctx = trace.ContextWithSpan(ctx, span)

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (b *Bcrypt) GenerateFromPassword(ctx context.Context, password string) (string, error) {
	ctx, span := b.tracer.Start(ctx, "bcryptRepo.CompareHashAndPassword")
	defer span.End()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
