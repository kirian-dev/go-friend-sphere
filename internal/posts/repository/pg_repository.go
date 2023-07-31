package usecase

import (
	"context"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/internal/posts"

	"github.com/jmoiron/sqlx"
)

type postsRepo struct {
	db *sqlx.DB
}

func NewPostsRepo(db *sqlx.DB) posts.Repository {
	return &postsRepo{db: db}
}

func (u *postsRepo) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	return nil, nil
}

func (u *postsRepo) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	return nil, nil
}

func (h *postsRepo) DeletePost(ctx context.Context, post *models.Post) error {
	return nil
}

func (h *postsRepo) GetPost(ctx context.Context, post *models.Post) (*models.Post, error) {
	return nil, nil
}

func (h *postsRepo) GetPosts(ctx context.Context, post *models.Post) (*models.Post, error) {
	return nil, nil
}
