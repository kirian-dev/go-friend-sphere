package models

import "time"

type Comment struct {
	CommentID int64     `json:"comment_id" db:"comment_id"`
	Message   string    `json:"message" db:"message" validate:"required,omitempty,lte=512"`
	UserID    int64     `json:"user_id" db:"user_id" validate:"required"`
	PostID    int64     `json:"post_id" db:"post_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CommentWithUser struct {
	CommentID int64     `json:"comment_id" db:"comment_id"`
	Message   string    `json:"message" db:"message" validate:"required,omitempty,lte=512"`
	UserID    int64     `json:"user_id" db:"user_id" validate:"required"`
	PostID    int64     `json:"post_id" db:"post_id" validate:"required"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
}
