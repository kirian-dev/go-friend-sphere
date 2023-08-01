package repository

import (
	"context"
	"database/sql"
	"go-friend-sphere/internal/messages"
	"go-friend-sphere/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type messagesRepo struct {
	db *sqlx.DB
}

func NewMessageRepo(db *sqlx.DB) messages.Repository {
	return &messagesRepo{db: db}
}

func (r *messagesRepo) CreateMessage(ctx context.Context, message *models.Message) (*models.Message, error) {
	m := &models.Message{}
	if err := r.db.QueryRowxContext(ctx, createMessage, message.Message, message.SenderID, message.RecipientID).StructScan(m); err != nil {
		return nil, errors.Wrap(err, "Message repository, Create Message")
	}

	return m, nil
}

func (r *messagesRepo) UpdateMessage(ctx context.Context, message *models.Message) (*models.Message, error) {
	updatedMessage := &models.Message{}

	if err := r.db.GetContext(ctx, updatedMessage, updateMessageQuery, &message.Message, &message.MessageID); err != nil {
		return nil, errors.Wrap(err, "Message repository, Update Message")
	}

	return updatedMessage, nil
}

func (r *messagesRepo) DeleteMessage(ctx context.Context, MessageID int64) error {
	result, err := r.db.ExecContext(ctx, deleteMessageQuery, MessageID)
	if err != nil {
		return errors.Wrap(err, "Message repository, Delete Message")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Message repository, RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "Message repository, rowsAffected")
	}

	return nil
}

func (r *messagesRepo) GetMessageByID(ctx context.Context, messageID int64) (*models.Message, error) {
	// ... (your implementation)
	// Handle the error from the scan operation
	message := &models.Message{}
	err := r.db.QueryRowContext(ctx, getMessageByID, messageID).Scan(message)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (r *messagesRepo) GetMessagesByUserID(ctx context.Context, userID int64) ([]*models.Message, error) {
	// ... (your implementation)
	// Handle the error from the rows.Scan operation
	messagesList := []*models.Message{}

	if err := r.db.SelectContext(ctx, &messagesList, getMessagesByUserID, userID); err != nil {
		return nil, errors.Wrap(err, "Comment repository, Get Comments")
	}

	return messagesList, nil
}
