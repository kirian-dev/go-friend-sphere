package models

import "time"

type User struct {
	UserID            int64     `json:"user_id"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Phone             *string   `json:"phone"`
	ProfilePictureUrl *string   `json:"profile_picture_url"`
	City              *string   `json:"city"`
	Birthday          *string   `json:"birthday"`
	UpdateAt          time.Time `json:"update_at"`
	CreatedAt         time.Time `json:"created_at"`
	LastLoginAt       time.Time `json:"last_login_at"`
}
