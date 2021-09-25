package models

import (
	"time"
)

type Commerce_Product struct {
	Product_ID       uint      `json:"product_id" gorm:"primary_key"`
	ProductName      string    `json:"product_name" gorm:"not null;unique"`
	ProductPict      string    `json:"product_pict"`
	ProductDesc      string    `json:"product_desc"`
	ProductPrice     int       `json:"product_price"`
	ProductCondition string    `json:"product_condition"`
	ProductWeight    int       `json:"product_weight"`
	ProductDimension string    `json:"product_dimension"`
	Etalase_ID       uint      `json:"etalase_id"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
