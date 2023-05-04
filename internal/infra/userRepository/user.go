package userRepository

import (
	"github.com/SafetyLink/authenticationService/internal/domain/repo"
	"github.com/SafetyLink/authenticationService/internal/infra/adapter/sql/sqlc"
	"github.com/jackc/pgx/v5"
)

type userRepository struct {
	database *sqlc.Queries
	db       *pgx.Conn
}

func NewUserRepository(database *pgx.Conn) repo.User {
	return &userRepository{
		database: sqlc.New(database),
		db:       database,
	}
}
