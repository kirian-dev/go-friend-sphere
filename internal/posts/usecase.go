package posts

import (
	"context"
	"go-friend-sphere/internal/models"
)

type UseCase interface {
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	DeletePost(ctx context.Context, postId int64) error
	GetPostById(ctx context.Context, postId int64) (*models.Post, error)
	GetPosts(ctx context.Context, params models.GetPostsParams) ([]*models.Post, error)
	ToggleLikePost(ctx context.Context, postId, userId int64) (bool, error)
}
