package models

import (
	"time"
)

type Commerce_Categories struct {
	Category_ID        uint      `json:"category_id" gorm:"primary_key"`
	Category_Name      string    `json:"category_name"`
	Category_Parent_ID uint      `json:"category_parent_id"`
	Category_Child_ID  uint      `json:"category_child_id"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
