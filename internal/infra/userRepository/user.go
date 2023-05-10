package userRepository

import (
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql/sqlc"
	"github.com/jackc/pgx/v5"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type userRepository struct {
	database *sqlc.Queries
	db       *pgx.Conn
	tracer   trace.Tracer
	logger   *zap.Logger
}

func NewUserRepository(database *pgx.Conn, logger *zap.Logger, tracer trace.Tracer) repo.User {

	return &userRepository{
		database: sqlc.New(database),
		db:       database,
		logger:   logger,
		tracer:   tracer,
	}
}
