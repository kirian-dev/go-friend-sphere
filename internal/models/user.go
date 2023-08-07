package models

import (
	"time"
)

type User struct {
	UserID            int64     `json:"user_id" db:"user_id"`
	FirstName         string    `json:"first_name" db:"first_name"`
	LastName          string    `json:"last_name" db:"last_name"`
	Email             string    `json:"email" db:"email"`
	Password          string    `json:"password" db:"password"`
	Role              string    `json:"role" db:"role"`
	Phone             *string   `json:"phone" db:"phone"`
	ProfilePictureUrl *string   `json:"profile_picture_url" db:"profile_picture_url"`
	City              *string   `json:"city" db:"city"`
	Birthday          *string   `json:"birthday" db:"birthday"`
	UpdateAt          time.Time `json:"update_at" db:"updated_at"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	LastLoginAt       time.Time `json:"last_login_at" db:"last_login_at"`
}
