package models

import (
	"database/sql/driver"
	"time"

	"gorm.io/gorm"
)

type boardPosition string

const (
	wishlist  boardPosition = "wishlist"
	applied   boardPosition = "applied"
	interview boardPosition = "interview"
	offer     boardPosition = "offer"
	rejected  boardPosition = "rejected"
)

func (self *boardPosition) Scan(value interface{}) error {
	*self = boardPosition(value.([]byte))
	return nil
}

func (self boardPosition) Value() (driver.Value, error) {
	return string(self), nil
}

type BoardStruct struct {
	gorm.Model
	BoardPosition boardPosition `gorm:"type:board_position"`
}

func (BoardStruct) TableName() string {
	return "Board"
}

type Job struct {
	gorm.Model
	ID            uint          `json:"id gorm:"primaryKey"`
	status        string        `json:"status"`
	CompanyName   string        `json:"companyName"`
	JobTitle      string        `json:"jobTitle"`
	BoardPosition boardPosition `gorm:"type:enum('wishlist', 'applied', 'interview', 'offer', 'rejected')";"column:board_position"`
	CreatedAt     time.Time     `json:"createdAt"`
	UpdatedAt     time.Time     `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt
	UserId        uint
}
