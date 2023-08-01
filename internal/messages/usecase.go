package messages

import (
	"context"
	"go-friend-sphere/internal/models"
)

type UseCase interface {
	CreateMessage(ctx context.Context, message *models.Message) (*models.Message, error)
	UpdateMessage(ctx context.Context, message *models.Message) (*models.Message, error)
	DeleteMessage(ctx context.Context, messageID int64) error
	GetMessageByID(ctx context.Context, messageID int64) (*models.Message, error)
	GetMessagesByUserID(ctx context.Context, userID int64) ([]*models.Message, error)
}
