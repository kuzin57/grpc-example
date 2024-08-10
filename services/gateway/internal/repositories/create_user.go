package repositories

import (
	"context"

	"github.com/kuzin57/grpc-example/services/gateway/internal/entities"
)

func (r *repository) CreateUser(ctx context.Context, user *entities.User) (string, error) {
	return "", nil
}
