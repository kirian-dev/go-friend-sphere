package models

import "time"

type Comment struct {
	CommentID int64     `json:"comment_id" db:"comment_id"`
	Message   string    `json:"message" db:"message"`
	UserID    int64     `json:"user_id" db:"user_id"`
	PostID    int64     `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"updated_at" db:"updated_at"`
}

type CommentWithUser struct {
	CommentID int64     `json:"comment_id" db:"comment_id"`
	Message   string    `json:"message" db:"message"`
	UserID    int64     `json:"user_id" db:"user_id"`
	PostID    int64     `json:"post_id" db:"post_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdateAt  time.Time `json:"updated_at" db:"updated_at"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
}
