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

func (u *postsUC) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	createdPost, err := u.postsRepo.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return createdPost, nil
}

func (u *postsUC) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	updatedPos, err := u.postsRepo.UpdatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return updatedPos, nil
}

func (u *postsUC) DeletePost(ctx context.Context, postId int64) error {
	return u.postsRepo.DeletePost(ctx, postId)

}

func (u *postsUC) GetPostById(ctx context.Context, postId int64) (*models.Post, error) {
	return u.postsRepo.GetPostById(ctx, postId)
}

func (u *postsUC) GetPosts(ctx context.Context, params models.GetPostsParams) ([]*models.Post, error) {
	return u.postsRepo.GetPosts(ctx, params)
}

func (u *postsUC) ToggleLikePost(ctx context.Context, postId, userId int64) (bool, error) {
	hasLiked, err := u.postsRepo.HasLikedPost(ctx, postId, userId)
	if err != nil {
		return true, err
	}

	if hasLiked {
		err = u.postsRepo.RemoveLike(ctx, postId, userId)
		if err != nil {
			return false, err
		}

		err = u.postsRepo.UpdateLikesCount(ctx, postId, -1)
		if err != nil {
			return false, err
		}
	} else {
		err = u.postsRepo.LikePost(ctx, postId, userId)
		if err != nil {
			return true, err
		}

		err = u.postsRepo.UpdateLikesCount(ctx, postId, +1)
		if err != nil {
			return true, err
		}
	}

	return true, nil
}
