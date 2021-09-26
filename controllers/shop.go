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

// GetShops godoc
// @Summary Get all shops.
// @Description Get a list of shops registered in the system.
// @Tags admin
// @Produce json
// @Security BearerToken
// @param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} []models.Commerce_Shop
// @Router /api/admin/shops [get]
func GetShops(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var shops []models.Commerce_Shop
	db.Find(&shops)

	c.JSON(http.StatusOK, gin.H{"data": shops})
}

// GetUserShops godoc
// @Summary Get user shop.
// @Description Get a list user shops in the system.
// @Tags user
// @Produce json
// @Success 200 {object} []models.Commerce_Shop
// @Router /api/user/shops [get]
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

// GetShopByKeyword godoc
// @Summary Search shop by keyword.
// @Description Get a list of shops registered in the system by keyword.
// @Tags public
// @Produce json
// @Param keyword query string true "The keyword for shop."
// @Success 200 {object} []models.Commerce_Shop
// @Router /api/shops/search [get]
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

// CreateShop godoc
// @Summary Create a shop.
// @Description Creating a shop from admin access.
// @Tags admin
// @Param Body body CreateShopInput true "the body to create a shop"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Commerce_Shop
// @Router /api/admin/shops/create [post]
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

// CreateShopUser godoc
// @Summary Create a shop by user.
// @Description Creating a shop from user access.
// @Tags user
// @Param Body body CreateShopInput true "the body to create a shop"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Commerce_Shop
// @Router /api/user/shops/create [post]
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

// FindShop godoc
// @Summary Get one shop.
// @Description Get a shop by its id.
// @Tags admin
// @Produce json
// @Security BearerToken
// @Param shop_id path string true "The shop id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Shop
// @Router /api/admin/shops/:shop_id [get]
func FindShop(c *gin.Context) { // Get model if exist
	var shop models.Commerce_Shop

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("shop_id = ?", c.Param("shop_id")).First(&shop).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": shop})
}

// ShopDetailUser godoc
// @Summary Get detail user shop.
// @Description Get a user shop detail.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Shop
// @Router /api/user/shop/:shop_id [get]
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

// UpdateShopUser godoc
// @Summary Update one user shop.
// @Description Update a user shop by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param shop_id path string true "The shop id"
// @Param Body body UpdateShopInput true "the body to update a shop"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Shop
// @Router /api/user/shop/:shop_id [patch]
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

// UpdateShop godoc
// @Summary Update one shop.
// @Description Update a shop by its id.
// @Tags admin
// @Produce json
// @Security BearerToken
// @Param shop_id path string true "The shop id"
// @Param Body body UpdateShopInput true "the body to update a shop"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Shop
// @Router /api/admin/shops/:shop_id [patch]
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

// DeleteShop godoc
// @Summary Delete one shop.
// @Description Delete a shop by its id.
// @Tags admin
// @Produce json
// @Security BearerToken
// @Param shop_id path string true "The shop id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/admin/shops/:shop_id [delete]
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

// DeleteShopUser godoc
// @Summary Delete one user shop.
// @Description Delete a shop user by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param shop_id path string true "The shop id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/user/shop/:shop_id [delete]
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
