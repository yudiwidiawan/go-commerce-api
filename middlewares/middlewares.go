package middlewares

import (
	"net/http"

	"fp-jcc-go-2021-commerce/models"
	"fp-jcc-go-2021-commerce/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func JwtAuthMiddlewareAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		user_id, err := token.ExtractTokenID(c)
		var user models.Commerce_User

		db := c.MustGet("db").(*gorm.DB)
		if err := db.Where("user_id = ?", user_id).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			c.Abort()
			return
		}
		if user.Role_ID != 1 {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		user_id, err := token.ExtractTokenID(c)
		var user models.Commerce_User
		db := c.MustGet("db").(*gorm.DB)
		if err := db.Where("user_id = ?", user_id).First(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			c.Abort()
			return
		}
		if user.Role_ID != 2 {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
