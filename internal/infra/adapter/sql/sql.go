package sql

import (
	"context"
	"github.com/SafetyLink/authenticationService/internal"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func NewSqlProvider(logger *zap.Logger, config *internal.Config) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config.Postgres.ConnectionURL)
	if err != nil {
		logger.Panic("failed to connect to postgres")
	}

	logger.Info("connected to postgres")

	return conn
}
