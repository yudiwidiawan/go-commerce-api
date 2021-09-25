package controllers

import (
	"net/http"

	"fp-jcc-go-2021-commerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateCategoryInput struct {
	Category_Name      string `json:"category_name"`
	Category_Parent_ID uint   `json:"category_parent_id"`
	Category_Child_ID  uint   `json:"category_child_id"`
}

type UpdateCategoryInput struct {
	Category_Name      string `json:"category_name"`
	Category_Parent_ID uint   `json:"category_parent_id"`
	Category_Child_ID  uint   `json:"category_child_id"`
}

// GET /categories
// Get all categories
func GetCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var categories []models.Commerce_Categories
	if err := db.Where("category_parent_id = 0").Find(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// POST /categories/create
// Create new category
func CreateCategory(c *gin.Context) {
	// Validate input
	var input CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create category
	category := models.Commerce_Categories{Category_Name: input.Category_Name, Category_Parent_ID: input.Category_Parent_ID,
		Category_Child_ID: input.Category_Child_ID}

	db := c.MustGet("db").(*gorm.DB)
	creatingCategory := db.Create(&category)
	if creatingCategory.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingCategory.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// GET /category/:category_id
// Get category by id
func GetCategory(c *gin.Context) { // Get model if exist
	var category models.Commerce_Categories

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("category_id = ? ", c.Param("category_id")).Find(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// PATCH /category/:category_id
// Update an category
func UpdateCategory(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var category models.Commerce_Categories
	if err := db.Where("category_id = ? ", c.Param("category_id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_Categories
	updatedInput.Category_Name = input.Category_Name
	updatedInput.Category_Parent_ID = input.Category_Parent_ID
	updatedInput.Category_Child_ID = input.Category_Child_ID

	db.Model(&category).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// DELETE /category/:category
// Delete a category
func DeleteCategory(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var category models.Commerce_Categories
	if err := db.Where("category_id = ? ", c.Param("category_id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}
	if (category.Category_Parent_ID != 0) && (category.Category_Child_ID == 0) {
		db.Delete(&category)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": category})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
