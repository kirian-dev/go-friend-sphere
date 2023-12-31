package models

import "time"

type Friendship struct {
	FriendshipID int64     `json:"friendship_id" db:"friendship_id"`
	UserID       int64     `json:"user_id" db:"user_id" validate:"required"`
	FriendID     int64     `json:"friend_id" db:"friend_id" validate:"required"`
	Status       string    `json:"status" db:"status validate:required" validate:"required,omitempty,lte=24"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type FriendshipWithFriend struct {
	FriendshipID int64     `json:"friendship_id" db:"friendship_id"`
	UserID       int64     `json:"user_id" db:"user_id" validate:"required"`
	FriendID     int64     `json:"friend_id" db:"friend_id" validate:"required"`
	Status       string    `json:"status" db:"status  validate:required,omitempty,lte=24"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	FirstName    string    `json:"first_name" db:"first_name"`
	LastName     string    `json:"last_name" db:"last_name"`
}
