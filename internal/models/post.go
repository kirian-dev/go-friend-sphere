package models

import "time"

type Post struct {
	PostId      int64     `json: "post_id db: post_id"`
	Content     string    `json: "content"`
	UserId      int64     `json: "user_id" db: user_id"`
	UpdateAt    time.Time `json: "update_at" db: update_at`
	created_at  time.Time `json: "created_at" db:created_at"`
	likes_count int64     `json: "likes_count" db: likes_count"`
	image_url   string    `json: "image_url" db: image_url"`
}
