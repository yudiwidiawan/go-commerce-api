package controllers

import (
	"net/http"

	"fp-jcc-go-2021-commerce/models"
	"fp-jcc-go-2021-commerce/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateCartInput struct {
	User_ID uint   `json:"user_id"`
	Status  string `json:"status"`
}

type UpdateCartInput struct {
	User_ID uint   `json:"user_id"`
	Status  string `json:"status"`
}

type Commerce_Cart_Product struct {
	Cart_Product_ID uint `json:"cart_product_id" gorm:"primary_key"`
	Cart_ID         uint `json:"cart_id"`
	Product_ID      uint `json:"product_id"`
	Product_Count   int  `json:"product_count"`
}

// GET /cart_items
// Get all cart items
// func GetCartItems(c *gin.Context) {
// 	db := c.MustGet("db").(*gorm.DB)

// 	var categories []models.Commerce_Product
// 	if err := db.Raw("SELECT commerce_products.*, commerce_cart_products.product_count FROM commerce_products, "+
// 	"").Find(&categories).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": categories})
// }

// POST /cart/create
// Create new cart
func CreateCart(c *gin.Context) {
	// Validate input
	var input CreateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user_id, _ := token.ExtractTokenID(c)
	// Create cart
	cart := models.Commerce_Cart_User{User_ID: user_id, Status: input.Status}

	db := c.MustGet("db").(*gorm.DB)
	creatingCart := db.Create(&cart)
	if creatingCart.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingCart.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// POST /cart/add
// Create new cart
func AddProductCart(c *gin.Context) {
	// Validate input
	var input Commerce_Cart_Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create cart
	cartProduct := Commerce_Cart_Product{Cart_ID: input.Cart_ID, Product_ID: input.Product_ID,
		Product_Count: input.Product_Count}

	db := c.MustGet("db").(*gorm.DB)
	addingCartProduct := db.Create(&cartProduct)
	if addingCartProduct.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": addingCartProduct.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cartProduct})
}

// PATCH /cart/:cart_id
// Update a cart
func UpdateCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var cart models.Commerce_Cart_User
	if err := db.Where("cart_id = ? ", c.Param("cart_id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateCartInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_Cart_User
	updatedInput.Status = input.Status

	db.Model(&cart).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": cart})
}

// PATCH /cart/product/:cart_product_id
// Update product cart
func UpdateProductCart(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var productCart Commerce_Cart_Product
	if err := db.Where("cart_product_id = ? ", c.Param("cart_product_id")).First(&productCart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input Commerce_Cart_Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&productCart).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": productCart})
}

// DELETE /cart/:cart_id
// Delete a cart
func DeleteCart(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var cart models.Commerce_Cart_User
	var cartProduct Commerce_Cart_Product
	if err := db.Where("cart_id = ? ", c.Param("cart_id")).First(&cartProduct).Error; err == nil {
		db.Delete(&cartProduct)
	}
	if err := db.Where("cart_id = ? ", c.Param("cart_id")).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}
	db.Delete(&cart)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DELETE /cart/product/:cart_product_id
// Delete a cart
func DeleteProductCart(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var cartProduct Commerce_Cart_Product
	if err := db.Where("cart_product_id = ? ", c.Param("cart_product_id")).First(&cartProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}
	db.Delete(&cartProduct)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
