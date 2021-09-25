package models

import (
	"time"
)

type Commerce_Cart_User struct {
	Cart_ID   uint      `json:"cart_id" gorm:"primary_key"`
	User_ID   uint      `json:"user_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
