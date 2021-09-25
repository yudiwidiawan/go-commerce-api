package controllers

import (
	"net/http"

	"fp-jcc-go-2021-commerce/models"
	"fp-jcc-go-2021-commerce/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateShopInput struct {
	User_ID         uint   `json:"user_id"`
	ShopName        string `json:"shop_name" gorm:"not null;unique"`
	ShopDesc        string `json:"shop_desc" gorm:"not null"`
	ShopEmail       string `json:"shop_email"`
	ShopPhoneNumber string `json:"shop_phone_number"`
}

type UpdateShopInput struct {
	User_ID         uint   `json:"user_id"`
	ShopName        string `json:"shop_name" gorm:"not null;unique"`
	ShopDesc        string `json:"shop_desc" gorm:"not null"`
	ShopEmail       string `json:"shop_email"`
	ShopPhoneNumber string `json:"shop_phone_number"`
}

// GET /shops
// Get all shops
func GetShops(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var shops []models.Commerce_Shop
	db.Find(&shops)

	c.JSON(http.StatusOK, gin.H{"data": shops})
}

func GetUserShops(c *gin.Context) {
	user_id, _ := token.ExtractTokenID(c)
	db := c.MustGet("db").(*gorm.DB)
	var shops []models.Commerce_Shop
	if err := db.Where("user_id = ?", user_id).Find(&shops).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shops})
}

// /shop/search?:keyword
// Search shop by keyword
func GetShopByKeyword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var shops []models.Commerce_Shop
	if err := db.Where("shop_name LIKE '%" + c.Query("keyword") + "%' OR shop_desc LIKE '%" + c.Query("keyword") + "%' ").Find(&shops).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shops})
}

// POST /shops/create
// Create new shop
func CreateShop(c *gin.Context) {
	// Validate input
	var input CreateShopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create shop
	shop := models.Commerce_Shop{User_ID: input.User_ID, ShopName: input.ShopName, ShopDesc: input.ShopDesc,
		ShopEmail: input.ShopEmail, ShopPhoneNumber: input.ShopPhoneNumber}

	db := c.MustGet("db").(*gorm.DB)
	creatingShop := db.Create(&shop)
	if creatingShop.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingShop.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

func CreateShopUser(c *gin.Context) {
	// Validate input
	var input CreateShopInput
	user_id, _ := token.ExtractTokenID(c)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create shop
	shop := models.Commerce_Shop{User_ID: user_id, ShopName: input.ShopName, ShopDesc: input.ShopDesc,
		ShopEmail: input.ShopEmail, ShopPhoneNumber: input.ShopPhoneNumber}

	db := c.MustGet("db").(*gorm.DB)
	creatingShop := db.Create(&shop)
	if creatingShop.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingShop.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// GET /shops/:shop_id
// Find a shop
func FindShop(c *gin.Context) { // Get model if exist
	var shop models.Commerce_Shop

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("shop_id = ?", c.Param("shop_id")).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// GET /details
// Get Self User Shop Detail
func ShopDetailUser(c *gin.Context) { // Get model if exist
	var shop models.Commerce_Shop
	user_id, err := token.ExtractTokenID(c)
	db := c.MustGet("db").(*gorm.DB)
	if err = db.Where("user_id = ? AND shop_id = ? ", user_id, c.Param("shop_id")).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// PATCH /update
// Update self shop
func UpdateShopUser(c *gin.Context) {
	user_id, _ := token.ExtractTokenID(c)
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var shop models.Commerce_Shop
	if err := db.Where("shop_id = ? AND user_id = ? ", c.Param("shop_id"), user_id).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateShopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_Shop
	updatedInput.ShopName = input.ShopName
	updatedInput.ShopDesc = input.ShopDesc
	updatedInput.ShopEmail = input.ShopEmail
	updatedInput.ShopPhoneNumber = input.ShopPhoneNumber

	db.Model(&shop).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// PATCH /shops/:shop_id
// Update a shop
func UpdateShop(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var shop models.Commerce_Shop
	if err := db.Where("shop_id = ?", c.Param("shop_id")).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateShopInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_Shop
	updatedInput.ShopName = input.ShopName
	updatedInput.ShopDesc = input.ShopDesc
	updatedInput.ShopEmail = input.ShopEmail
	updatedInput.ShopPhoneNumber = input.ShopPhoneNumber

	db.Model(&shop).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// DELETE /shops/:id
// Delete a shop
func DeleteShop(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var shop models.Commerce_Shop
	if err := db.Where("shop_id = ?", c.Param("shop_id")).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&shop)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func DeleteShopUser(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	user_id, _ := token.ExtractTokenID(c)
	var shop models.Commerce_Shop
	if err := db.Where("shop_id = ? and user_id = ? ", c.Param("shop_id"), user_id).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&shop)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
