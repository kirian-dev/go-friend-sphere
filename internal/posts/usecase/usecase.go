package usecase

import (
	"context"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/internal/posts"
	"go-friend-sphere/pkg/logger"
)

type postsUC struct {
	cfg       *config.Config
	logger    logger.ZapLogger
	postsRepo posts.Repository
}

func NewPostsUC(cfg *config.Config, logger logger.ZapLogger, postsRepo posts.Repository) posts.UseCase {
	return &postsUC{cfg: cfg, logger: logger, postsRepo: postsRepo}
}

func (h *postsUC) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	return nil, nil
}

func (h *postsUC) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	return nil, nil
}

func (h *postsUC) DeletePost(ctx context.Context, post *models.Post) error {
	return nil
}

func (h *postsUC) GetPost(ctx context.Context, post *models.Post) (*models.Post, error) {
	return nil, nil
}

func (h *postsUC) GetPosts(ctx context.Context) ([]*models.Post, error) {
	return nil, nil
}
