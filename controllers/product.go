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

// GET /shop/:shop_id/products
// Get all products in shop
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

// /products/search?:keyword
// Search products by keyword
func GetProductsByKeyword(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Commerce_Product
	if err := db.Where("product_name LIKE '%" + c.Query("keyword") + "%' OR product_desc LIKE '%" + c.Query("keyword") + "%' ").Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// /products/:product_id/categories
// Get products categories
func GetProductCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var productCategories []Commerce_Product_Categories
	if err := db.Where("product_id = ? ", c.Param("product_id")).Find(&productCategories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": productCategories})
}

// POST /etalase/:etalase_id/product/create
// Create new products
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

// POST /product/category/create
// Add Product Category
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

// GET /etalase/:etalase_id/products
// Get all products in an etalase
func GetProducts(c *gin.Context) { // Get model if exist
	var products []models.Commerce_Product

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("etalase_id = ? ", c.Param("etalase_id")).Find(&products).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})
}

// PATCH /product/:product_id
// Update a product
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

// PATCH /product/category/:product_category_id
// Update a product category
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

// DELETE /product/:id_product
// Delete an etalase in a shop
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

// DELETE /product/category/:product_category_id
// Delete a product category
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
