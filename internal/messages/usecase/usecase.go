package usecase

import (
	"context"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/friendships"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/logger"
)

type friendshipsUC struct {
	cfg             *config.Config
	logger          logger.ZapLogger
	friendshipsRepo friendships.Repository
}

func NewFriendshipsUC(cfg *config.Config, logger logger.ZapLogger, friendshipsRepo friendships.Repository) friendships.UseCase {
	return &friendshipsUC{cfg: cfg, logger: logger, friendshipsRepo: friendshipsRepo}
}

func (u *friendshipsUC) CreateFriendship(ctx context.Context, friendship *models.Friendship) (*models.Friendship, error) {
	createdFriendship, err := u.friendshipsRepo.CreateFriendship(ctx, friendship)
	if err != nil {
		return nil, err
	}

	return createdFriendship, nil
}

func (u *friendshipsUC) UpdateFriendship(ctx context.Context, friendship *models.Friendship) (*models.Friendship, error) {
	updateF, err := u.friendshipsRepo.UpdateFriendship(ctx, friendship)
	if err != nil {
		return nil, err
	}

	return updateF, nil
}

func (u *friendshipsUC) DeleteFriendship(ctx context.Context, friendshipID int64) error {
	return u.friendshipsRepo.DeleteFriendship(ctx, friendshipID)

}

func (u *friendshipsUC) GetFriendshipByID(ctx context.Context, friendshipID int64) (*models.FriendshipWithFriend, error) {
	return u.friendshipsRepo.GetFriendshipByID(ctx, friendshipID)
}

func (u *friendshipsUC) GetFriendshipsByUserID(ctx context.Context, postID int64) ([]*models.FriendshipWithFriend, error) {
	return u.friendshipsRepo.GetFriendshipsByUserID(ctx, postID)
}
