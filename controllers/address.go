package controllers

import (
	"net/http"

	"fp-jcc-go-2021-commerce/models"
	"fp-jcc-go-2021-commerce/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateAddressInput struct {
	AddressStreet    string `json:"address_street"`
	AddressProvince  string `json:"address_province"`
	AddressCity      string `json:"address_city"`
	AddressCountry   string `json:"address_country"`
	AddressPostcode  int    `json:"address_postcode"`
	AddressLatitude  string `json:"address_latitude"`
	AddressLongitude string `json:"address_longitude"`
}

type CommerceUserAddress struct {
	User_ID        uint
	Address_ID     uint
	Address_Status string
}

type UpdateAddressInput struct {
	AddressStreet    string `json:"address_street"`
	AddressProvince  string `json:"address_province"`
	AddressCity      string `json:"address_city"`
	AddressCountry   string `json:"address_country"`
	AddressPostcode  int    `json:"address_postcode"`
	AddressLatitude  string `json:"address_latitude"`
	AddressLongitude string `json:"address_longitude"`
}

// GetAddresses godoc
// @Summary Get all address from user.
// @Description Get a list of available address user.
// @Tags user
// @Produce json
// @Success 200 {object} []models.Commerce_Address
// @Router /api/user/addresses [get]
func GetAddresses(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	user_id, _ := token.ExtractTokenID(c)
	var addresses []models.Commerce_Address
	if err := db.Raw("SELECT commerce_addresses.* FROM commerce_addresses "+
		"WHERE commerce_addresses.address_id IN (SELECT address_id FROM commerce_user_addresses "+
		"WHERE commerce_user_addresses.user_id = ? )", user_id).Scan(&addresses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": addresses})
}

// CreateAddress godoc
// @Summary Create an address.
// @Description Creating a new address user.
// @Tags user
// @Param Body body CreateAddressInput true "the body to create a new address"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Commerce_Address
// @Router /api/user/address/create [post]
func CreateAddress(c *gin.Context) {
	// Validate input
	var input CreateAddressInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// user_id, _ := token.ExtractTokenID(c)

	// Create address
	address := models.Commerce_Address{AddressStreet: input.AddressStreet, AddressProvince: input.AddressProvince,
		AddressCity: input.AddressCity, AddressCountry: input.AddressCountry, AddressPostcode: input.AddressPostcode,
		AddressLatitude: input.AddressLatitude, AddressLongitude: input.AddressLongitude}

	db := c.MustGet("db").(*gorm.DB)
	creatingAddress := db.Create(&address)
	if creatingAddress.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingAddress.Error})
		return
	}
	user_id, _ := token.ExtractTokenID(c)

	addressUser := CommerceUserAddress{
		User_ID:        uint(user_id),
		Address_ID:     address.Address_ID,
		Address_Status: "Not Primary",
	}
	creatingAddressUser := db.Create(&addressUser)
	if creatingAddressUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingAddressUser.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": address})
}

// GET /address/:address_id
// Get all products in an etalase
func GetAddress(c *gin.Context) { // Get model if exist
	var address []models.Commerce_Address

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("address_id = ? ", c.Param("address_id")).Find(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": address})
}

// UpdateAddress godoc
// @Summary Update an address.
// @Description Update an address by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param address_id path string true "The address id"
// @Param Body body UpdateAddressInput true "the body to update an address"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Address
// @Router /api/user/address/:address_id [patch]
func UpdateAddress(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var address models.Commerce_Address
	if err := db.Where("address_id = ? ", c.Param("address_id")).First(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateAddressInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_Address
	updatedInput.AddressStreet = input.AddressStreet
	updatedInput.AddressProvince = input.AddressProvince
	updatedInput.AddressCity = input.AddressCity
	updatedInput.AddressCountry = input.AddressCountry
	updatedInput.AddressPostcode = input.AddressPostcode
	updatedInput.AddressLatitude = input.AddressLatitude
	updatedInput.AddressLongitude = input.AddressLongitude

	db.Model(&address).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": address})
}

// DeleteAddress godoc
// @Summary Delete one address.
// @Description Delete an address by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param address_id path string true "The address id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/user/address/:address_id [delete]
func DeleteAddress(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var address models.Commerce_Address
	var addressUser CommerceUserAddress
	if err := db.Where("address_id = ? ", c.Param("address_id")).First(&addressUser).Error; err == nil {
		db.Delete(&addressUser)
	}

	if err := db.Where("address_id = ? ", c.Param("address_id")).First(&address).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}
	db.Delete(&address)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
