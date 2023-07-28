package auth

import (
	"context"
	"go-friend-sphere/internal/models"
)

type UseCase interface {
	Register(ctx context.Context, user *models.User) (*models.User, error)
	GetUsers(ctx context.Context) ([]*models.User, error)
	Login(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, userId int64) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, userId int64) error
}
