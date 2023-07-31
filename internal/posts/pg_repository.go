package posts

import (
	"context"
	"go-friend-sphere/internal/models"
)

type Repository interface {
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	DeletePost(ctx context.Context, postId int64) error
	GetPostById(ctx context.Context, postId int64) (*models.Post, error)
	GetPosts(ctx context.Context) ([]*models.Post, error)
	HasLikedPost(ctx context.Context, postId, userId int64) (bool, error)
	LikePost(ctx context.Context, postId, userId int64) error
	RemoveLike(ctx context.Context, postId, userId int64) error
	UpdateLikesCount(ctx context.Context, postId int64, count int) error
}
