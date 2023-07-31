package posts

import (
	"context"
	"go-friend-sphere/internal/models"
)

type Repository interface {
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	DeletePost(ctx context.Context, post *models.Post) error
	GetPost(ctx context.Context, post *models.Post) (*models.Post, error)
	GetPosts(ctx context.Context) ([]*models.Post, error)
}
