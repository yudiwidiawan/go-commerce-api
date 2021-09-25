package models

import (
	"time"
)

type Commerce_Shop struct {
	Shop_ID         uint      `json:"shop_id" gorm:"primary_key"`
	User_ID         uint      `json:"user_id"`
	ShopName        string    `json:"shop_name" gorm:"not null;unique"`
	ShopDesc        string    `json:"shop_desc" gorm:"not null"`
	ShopEmail       string    `json:"shop_email"`
	ShopPhoneNumber string    `json:"shop_phone_number"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
