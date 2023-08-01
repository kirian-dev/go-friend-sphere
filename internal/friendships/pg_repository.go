package comments

import (
	"context"
	"go-friend-sphere/internal/models"
)

type Repository interface {
	CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error)
	DeleteComment(ctx context.Context, commentID int64) error
	GetCommentByID(ctx context.Context, commentID int64) (*models.CommentWithUser, error)
	GetCommentsByPostID(ctx context.Context, postId int64) ([]*models.CommentWithUser, error)
}
