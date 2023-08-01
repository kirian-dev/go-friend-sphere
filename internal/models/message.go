package models

import "time"

type Message struct {
	MessageID   int64     `json:"message_id" db:"message_id"`
	Message     string    `json:"message" db:"message"`
	SenderID    int64     `json:"sender_id" db:"sender_id"`
	RecipientID int64     `json:"recipient_id" db:"recipient_id"`
	SentAt      time.Time `json:"sent_at" db:"sent_at"`
	ReadAt      time.Time `json:"read_at" db:"read_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
