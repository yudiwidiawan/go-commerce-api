package controllers

import (
	"net/http"
	"strconv"

	"fp-jcc-go-2021-commerce/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateEtalaseInput struct {
	EtalaseName string `json:"etalase_name" gorm:"not null;unique"`
	Shop_ID     uint   `json:"shop_id"`
}

type UpdateEtalaseInput struct {
	EtalaseName string `json:"etalase_name" gorm:"not null;unique"`
	Shop_ID     uint   `json:"shop_id"`
}

// GET /shop/:shop_id/etalase/all
// Get all etalase
func GetEtalases(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var etalases []models.Commerce_Etalase
	if err := db.Where("shop_id = ?", c.Param("shop_id")).Find(&etalases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": etalases})
}

// POST /shops/:shop_id/etalase/create
// Create new etalase
func CreateEtalase(c *gin.Context) {
	// Validate input
	var input CreateEtalaseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shop_id, _ := strconv.Atoi(c.Param("shop_id"))
	// Create etalase
	etalase := models.Commerce_Etalase{EtalaseName: input.EtalaseName, Shop_ID: uint(shop_id)}

	db := c.MustGet("db").(*gorm.DB)
	creatingEtalase := db.Create(&etalase)
	if creatingEtalase.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingEtalase.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": etalase})
}

// GET /shops/:shop_id/etalase/:etalase_id
// Find an etalase in a shop
func FindEtalase(c *gin.Context) { // Get model if exist
	var etalase models.Commerce_Etalase

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("shop_id = ? AND etalase_id = ? ", c.Param("shop_id"), c.Param("etalase_id")).First(&etalase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": etalase})
}

// PATCH /shops/:shop_id/etalase/:etalase_id
// Update self shop
func UpdateEtalase(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var etalase models.Commerce_Etalase
	if err := db.Where("shop_id = ? AND etalase_id = ? ", c.Param("shop_id"), c.Param("etalase_id")).First(&etalase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateEtalaseInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_Etalase
	updatedInput.EtalaseName = input.EtalaseName
	updatedInput.Shop_ID = input.Shop_ID

	db.Model(&etalase).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": etalase})
}

// DELETE /shop/:shop_id/etalase/:etalase_id
// Delete an etalase in a shop
func DeleteEtalase(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var etalase models.Commerce_Etalase
	if err := db.Where("shop_id = ? AND etalase_id = ? ", c.Param("shop_id"), c.Param("etalase_id")).First(&etalase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&etalase)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
