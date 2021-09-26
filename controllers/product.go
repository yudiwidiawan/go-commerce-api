package controllers

import (
	"net/http"
	"strconv"

	"fp-jcc-go-2021-commerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateProductInput struct {
	ProductName      string `json:"product_name" gorm:"not null;unique"`
	ProductPict      string `json:"product_pict"`
	ProductDesc      string `json:"product_desc"`
	ProductPrice     int    `json:"product_price"`
	ProductCondition string `json:"product_condition"`
	ProductWeight    int    `json:"product_weight"`
	ProductDimension string `json:"product_dimension"`
	Etalase_ID       uint   `json:"etalase_id"`
}

type UpdateProductInput struct {
	ProductName      string `json:"product_name" gorm:"not null;unique"`
	ProductPict      string `json:"product_pict"`
	ProductDesc      string `json:"product_desc"`
	ProductPrice     int    `json:"product_price"`
	ProductCondition string `json:"product_condition"`
	ProductWeight    int    `json:"product_weight"`
	ProductDimension string `json:"product_dimension"`
}

type Commerce_Product_Categories struct {
	Product_Category_ID uint `json:"product_category_id"`
	Product_ID          uint `json:"product_id"`
	Category_ID         uint `json:"category_id"`
}

type UpdateProductCategories struct {
	Product_Category_ID uint `json:"product_category_id"`
	Product_ID          uint `json:"product_id"`
	Category_ID         uint `json:"category_id"`
}

// GetShopProducts godoc
// @Summary Get products by shop.
// @Description Get a list of products in shop by shop id.
// @Tags user
// @Produce json
// @Param shop_id path string true "The shop id"
// @Success 200 {object} []models.Commerce_Product
// @Router /api/user/shop/:shop_id/products [get]
func GetShopProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Commerce_Product
	if err := db.Raw("SELECT commerce_products.* FROM commerce_products "+
		"WHERE commerce_products.etalase_id IN (SELECT etalase_id FROM commerce_etalases "+
		"WHERE commerce_etalases.shop_id = ? )", c.Param("shop_id")).Scan(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GetProductsByKeyword godoc
// @Summary Search products by keyword.
// @Description Get a list of products registered in the system by keyword.
// @Tags public
// @Produce json
// @Param keyword query string true "The keyword for products."
// @Success 200 {object} []models.Commerce_Product
// @Router /api/products/search [get]
func GetProductsByKeyword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Commerce_Product
	if err := db.Where("product_name LIKE '%" + c.Query("keyword") + "%' OR product_desc LIKE '%" + c.Query("keyword") + "%' ").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// GetProductCategories godoc
// @Summary Get product categories..
// @Description Get a list of product categories by product id.
// @Tags user
// @Produce json
// @Param product_id path string true "The product id."
// @Success 200 {object} []Commerce_Product_Categories
// @Router /api/user/products/:product_id/categories [get]
func GetProductCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var productCategories []Commerce_Product_Categories
	if err := db.Where("product_id = ? ", c.Param("product_id")).Find(&productCategories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": productCategories})
}

// CreateProduct godoc
// @Summary Create a product by user.
// @Description Creating a product from user access.
// @Tags user
// @Param etalase_id path string true "The etalase id"
// @Param Body body CreateProductInput true "the body to create a shop"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Commerce_Product
// @Router /api/user/etalase/:etalase_id/product/create [post]
func CreateProduct(c *gin.Context) {
	// Validate input
	var input CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var etalase_id int
	if c.Param("etalase_id") != "" {
		etalase_id, _ = strconv.Atoi(c.Param("etalase_id"))
	} else {
		etalase_id = int(input.Etalase_ID)
	}

	// Create product
	product := models.Commerce_Product{ProductName: input.ProductName, ProductPict: input.ProductPict,
		ProductDesc: input.ProductDesc, ProductPrice: input.ProductPrice, ProductCondition: input.ProductCondition,
		ProductWeight: input.ProductWeight, ProductDimension: input.ProductDimension, Etalase_ID: uint(etalase_id)}

	db := c.MustGet("db").(*gorm.DB)
	creatingProduct := db.Create(&product)
	if creatingProduct.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingProduct.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// CreateProductCategory godoc
// @Summary Create a product category by user.
// @Description Creating a product category from user access.
// @Tags user
// @Param Body body Commerce_Product_Categories true "the body to create a product category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} Commerce_Product_Categories
// @Router /api/user/product/category/create [post]
func CreateProductCategory(c *gin.Context) {
	// Validate input
	var input Commerce_Product_Categories
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create product category
	productCategory := Commerce_Product_Categories{Product_ID: input.Product_ID, Category_ID: input.Category_ID}
	db := c.MustGet("db").(*gorm.DB)
	creatingProductCat := db.Create(&productCategory)
	if creatingProductCat.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingProductCat.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": productCategory})
}

// GetProducts godoc
// @Summary Get products by etalase id.
// @Description Get a list of products in etalase.
// @Tags user
// @Produce json
// @Param etalase_id path string true "The etalase id"
// @Success 200 {object} []models.Commerce_Product
// @Router /api/user/etalase/:etalase_id/products [get]
func GetProducts(c *gin.Context) { // Get model if exist
	var products []models.Commerce_Product

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("etalase_id = ? ", c.Param("etalase_id")).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// UpdateProduct godoc
// @Summary Update one product.
// @Description Update a product by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param product_id path string true "The product id"
// @Param Body body UpdateProductInput true "the body to update a product"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Product
// @Router /api/user/product/:product_id [patch]
func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Commerce_Product
	if err := db.Where("product_id = ? ", c.Param("product_id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_Product
	updatedInput.ProductName = input.ProductName
	updatedInput.ProductPict = input.ProductPict
	updatedInput.ProductDesc = input.ProductDesc
	updatedInput.ProductPrice = input.ProductPrice
	updatedInput.ProductCondition = input.ProductCondition
	updatedInput.ProductWeight = input.ProductWeight
	updatedInput.ProductDimension = input.ProductDimension

	db.Model(&product).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": product})
}

// UpdateProductCategory godoc
// @Summary Update one product category id.
// @Description Update a product category by product category id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param product_category_id path string true "The product category id"
// @Param Body body UpdateProductCategories true "the body to update a product category"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} UpdateProductCategories
// @Router /api/user/product/category/:product_category_id [patch]
func UpdateProductCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var productCategory Commerce_Product_Categories
	if err := db.Where("product_category_id = ? ", c.Param("product_category_id")).First(&productCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateProductCategories
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Exec("UPDATE commerce_product_categories SET product_id = ?, category_id = ? "+
		"WHERE product_category_id = ?", input.Product_ID, input.Category_ID,
		c.Param("product_category_id"))

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// DeleteProduct godoc
// @Summary Delete a product.
// @Description Delete a user product by product id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param product_id path string true "The product id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/user/product/:product_id [delete]
func DeleteProduct(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var product models.Commerce_Product
	if err := db.Where("product_id = ? ", c.Param("product_id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DeleteProductCategory godoc
// @Summary Delete a product category.
// @Description Delete a user product category by product category id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param product_category_id path string true "The product category id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/user/product/category/:product_category_id [delete]
func DeleteProductCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var productCategory Commerce_Product_Categories
	if err := db.Where("product_category_id = ? ", c.Param("product_category_id")).First(&productCategory).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Exec("DELETE FROM commerce_product_categories WHERE product_category_id = ?", c.Param("product_category_id"))

	c.JSON(http.StatusOK, gin.H{"data": true})
}
