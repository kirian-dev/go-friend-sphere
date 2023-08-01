package usecase

import (
	"context"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/comments"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/logger"
)

type commentsUC struct {
	cfg          *config.Config
	logger       logger.ZapLogger
	commentsRepo comments.Repository
}

func NewCommentsUC(cfg *config.Config, logger logger.ZapLogger, commentsRepo comments.Repository) comments.UseCase {
	return &commentsUC{cfg: cfg, logger: logger, commentsRepo: commentsRepo}
}

func (u *commentsUC) CreateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	createdComment, err := u.commentsRepo.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return createdComment, nil
}

func (u *commentsUC) UpdateComment(ctx context.Context, comment *models.Comment) (*models.Comment, error) {
	updateCom, err := u.commentsRepo.UpdateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return updateCom, nil
}

func (u *commentsUC) DeleteComment(ctx context.Context, commentID int64) error {
	return u.commentsRepo.DeleteComment(ctx, commentID)

}

func (u *commentsUC) GetCommentByID(ctx context.Context, commentID int64) (*models.CommentWithUser, error) {
	return u.commentsRepo.GetCommentByID(ctx, commentID)
}

func (u *commentsUC) GetCommentsByPostID(ctx context.Context, postID int64) ([]*models.CommentWithUser, error) {
	return u.commentsRepo.GetCommentsByPostID(ctx, postID)
}
