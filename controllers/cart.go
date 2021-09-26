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

// CreateCart godoc
// @Summary Create a cart.
// @Description Creating a new cart for user.
// @Tags user
// @Param Body body CreateCartInput true "the body to create a new cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Commerce_Cart_User
// @Router /api/user/cart/create [post]
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

// AddProductCart godoc
// @Summary Add a product to a cart.
// @Description Adding a new product to an existing cart for user.
// @Tags user
// @Param Body body Commerce_Cart_Product true "the body to add a product to cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} Commerce_Cart_Product
// @Router /api/user/cart/add [post]
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

// UpdateCart godoc
// @Summary Update a cart.
// @Description Update a cart by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param cart_id path string true "The cart id"
// @Param Body body UpdateCartInput true "the body to update a cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Cart_User
// @Router /api/user/cart/:cart_id [patch]
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

// UpdateProductCart godoc
// @Summary Update product in a cart.
// @Description Updating a product in a cart by cart product id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param cart_product_id path string true "The cart product id"
// @Param Body body Commerce_Cart_Product true "the body to update a product cart"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} Commerce_Cart_Product
// @Router /api/user/cart/product/:cart_product_id [patch]
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

// DeleteCart godoc
// @Summary Delete one cart.
// @Description Delete a cart by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param cart_id path string true "The cart id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/user/cart/:cart_id [delete]
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

// DeleteProductCart godoc
// @Summary Delete a product cart.
// @Description Delete a product cart by cart product id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param cart_product_id path string true "The cart product id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/user/cart/product/:cart_product_id [delete]
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
