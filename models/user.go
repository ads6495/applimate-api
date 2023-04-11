package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint      `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time `json:"createAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	FirstName    *string   `json:"firstName"`
	LastName     *string   `json:"lastName"`
	Email        *string   `json:"email"`
	Password     *string   `json:"password"`
	Phone        *string   `json:"phone"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refreshToken"`
	Job          []Job     `json:"job" bson:"job"`
}
