package friendships

import (
	"context"
	"go-friend-sphere/internal/models"
)

type Repository interface {
	CreateFriendship(ctx context.Context, friendship *models.Friendship) (*models.Friendship, error)
	UpdateFriendship(ctx context.Context, friendship *models.Friendship) (*models.Friendship, error)
	DeleteFriendship(ctx context.Context, friendshipID int64) error
	GetFriendshipByID(ctx context.Context, friendshipID int64) (*models.FriendshipWithFriend, error)
	GetFriendshipsByUserID(ctx context.Context, userID int64) ([]*models.FriendshipWithFriend, error)
}
