package repositories

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kuzin57/grpc-example/services/gateway/internal/entities"

	_ "github.com/lib/pq"
)

const (
	pgDriverName = "postgres"
)

type ServiceRepository interface {
	CreateUser(ctx context.Context, user *entities.User) (string, error)
}

type repository struct {
	db *sqlx.DB
}

func NewServiceRepository(dsn string) (ServiceRepository, error) {
	db, err := sqlx.Connect(pgDriverName, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return &repository{
		db: db,
	}, nil
}
