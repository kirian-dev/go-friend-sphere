package usecase

import (
	"context"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/messages"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/logger"
)

type messagesUC struct {
	cfg          *config.Config
	logger       logger.ZapLogger
	messagesRepo messages.Repository
}

func NewMessagesUC(cfg *config.Config, logger logger.ZapLogger, messagesRepo messages.Repository) messages.UseCase {
	return &messagesUC{cfg: cfg, logger: logger, messagesRepo: messagesRepo}
}

func (u *messagesUC) CreateMessage(ctx context.Context, message *models.Message) (*models.Message, error) {
	createdMessage, err := u.messagesRepo.CreateMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return createdMessage, nil
}

func (u *messagesUC) UpdateMessage(ctx context.Context, message *models.Message) (*models.Message, error) {
	updateF, err := u.messagesRepo.UpdateMessage(ctx, message)
	if err != nil {
		return nil, err
	}

	return updateF, nil
}

func (u *messagesUC) DeleteMessage(ctx context.Context, messageID int64) error {
	return u.messagesRepo.DeleteMessage(ctx, messageID)

}

func (u *messagesUC) GetMessageByID(ctx context.Context, messageID int64) (*models.Message, error) {
	message, err := u.messagesRepo.GetMessageByID(ctx, messageID)
	if err != nil {
		return nil, err
	}

	currentUserID := int64(456)

	if message.RecipientID == currentUserID && message.ReadAt.IsZero() {
		if err := u.messagesRepo.ReadMessage(ctx, messageID); err != nil {
			return nil, err
		}
	}

	return message, nil
}

func (u *messagesUC) GetMessagesByUserID(ctx context.Context, userID int64) ([]*models.Message, error) {
	return u.messagesRepo.GetMessagesByUserID(ctx, userID)
}
