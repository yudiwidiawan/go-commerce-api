package models

import (
	"time"
)

type Commerce_Address struct {
	Address_ID       uint      `json:"address_id" gorm:"primary_key"`
	AddressStreet    string    `json:"address_street"`
	AddressProvince  string    `json:"address_province"`
	AddressCity      string    `json:"address_city"`
	AddressCountry   string    `json:"address_country"`
	AddressPostcode  int       `json:"address_postcode"`
	AddressLatitude  string    `json:"address_latitude"`
	AddressLongitude string    `json:"address_longitude"`
	CreatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
