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

// GetEtalases godoc
// @Summary Get etalase of a shop user.
// @Description Get etalase shop user by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param shop_id path string true "The shop id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} []models.Commerce_Etalase
// @Router /api/user/shop/:shop_id/etalase [get]
func GetEtalases(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var etalases []models.Commerce_Etalase
	if err := db.Where("shop_id = ?", c.Param("shop_id")).Find(&etalases).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": etalases})
}

// CreateEtalase godoc
// @Summary Create an etalase shop user.
// @Description Creating an etalase shop user.
// @Tags user
// @Param Body body CreateEtalaseInput true "the body to create an etalase"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Commerce_Etalase
// @Router /api/user/shop/:shop_id/etalase/create [post]
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

// FindEtalase godoc
// @Summary Get one etalase from one shop.
// @Description Get a etalase shop by shop id and etalase id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param shop_id path string true "The shop id"
// @Param etalase_id path string true "The etalase id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Etalase
// @Router /api/user/shop/:shop_id/etalase/:etalase_id [get]
func FindEtalase(c *gin.Context) { // Get model if exist
	var etalase models.Commerce_Etalase

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("shop_id = ? AND etalase_id = ? ", c.Param("shop_id"), c.Param("etalase_id")).First(&etalase).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": etalase})
}

// UpdateEtalase godoc
// @Summary Update an user shop etalase.
// @Description Update a user shop etalase by shop id and etalase id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param etalase_id path string true "The etalase id"
// @Param shop_id path string true "The shop id"
// @Param Body body UpdateEtalaseInput true "the body to update an etalase"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} models.Commerce_Etalase
// @Router /api/user/shop/:shop_id/etalase/:etalase_id [patch]
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

// DeleteEtalase godoc
// @Summary Delete one user etalase.
// @Description Delete a shop user etalase by its id.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param etalase_id path string true "The etalase id"
// @Param shop_id path string true "The shop id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]boolean
// @Router /api/user/shop/:shop_id/etalase/:etalase_id [delete]
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
