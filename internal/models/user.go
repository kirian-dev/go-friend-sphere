package models

import (
	"time"
)

type User struct {
	UserID            int64     `json:"user_id" db:"user_id"`
	FirstName         string    `json:"first_name" db:"first_name" validate:"required, lte=30"`
	LastName          string    `json:"last_name" db:"last_name" validate:"required, lte=30"`
	Email             string    `json:"email" db:"email" validate:"required,email, lte=60, omitempty"`
	Password          string    `json:"password" db:"password" validate:"omitempty,required,gte=6"`
	Role              string    `json:"role" db:"role"`
	Phone             *string   `json:"phone" db:"phone" validate:"omitempty,lte=20"`
	ProfilePictureUrl *string   `json:"profile_picture_url" db:"profile_picture_url" validate:"omitempty,lte=512,url"`
	City              *string   `json:"city" db:"city" validate:"omitempty,lte=24"`
	Birthday          *string   `json:"birthday" db:"birthday" validate:"omitempty,lte=10"`
	UpdateAt          time.Time `json:"update_at" db:"updated_at"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	LastLoginAt       time.Time `json:"last_login_at" db:"last_login_at"`
}
