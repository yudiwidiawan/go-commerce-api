package models

import (
	"time"
)

type Commerce_User struct {
	User_ID     uint      `json:"user_id" gorm:"primary_key"`
	Username    string    `json:"username" gorm:"not null;unique"`
	Password    string    `json:"password" gorm:"not null"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Gender      string    `json:"gender"`
	ProfilePic  string    `json:"profile_pic"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Role_ID     uint
}
