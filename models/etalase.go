package models

import (
	"time"
)

type Commerce_Etalase struct {
	Etalase_ID  uint      `json:"etalase_id" gorm:"primary_key"`
	EtalaseName string    `json:"etalase_name" gorm:"not null;unique"`
	Shop_ID     uint      `json:"shop_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
