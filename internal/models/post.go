package models

import "time"

type Post struct {
	PostID     int64     `json:"post_id" db:"post_id"`
	Content    string    `json:"content"`
	UserId     int64     `json:"user_id" db:"user_id"`
	UpdateAt   time.Time `json:"updated_at" db:"updated_at"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	LikesCount int64     `json:"likes_count" db:"likes_count"`
	ImageUrl   string    `json:"image_url" db:"image_url"`
}
