package auth

import (
	"context"
	"go-friend-sphere/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
}
