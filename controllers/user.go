package controllers

import (
	"html"
	"net/http"
	"strings"

	"fp-jcc-go-2021-commerce/models"
	"fp-jcc-go-2021-commerce/utils/token"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type CreateUserInput struct {
	gorm.Model
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	ProfilePic  string `json:"profile_pic"`
}

type UpdateUserInput struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
	ProfilePic  string `json:"profile_pic"`
}

type ReturnSimpleUser struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// LOGIN /login
// Login user
func LoginUser(c *gin.Context) {

	var input LoginUserInput
	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Commerce_User{}

	u.Username = input.Username
	u.Password = input.Password

	err := db.Model(models.Commerce_User{}).Where("username = ?", input.Username).Take(&u).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(input.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error password": err.Error()})
		return
	}

	token, err := token.GenerateToken(u.User_ID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"user": u.User_ID, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.Commerce_User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.Commerce_User{Username: input.Username, Password: input.Password, FirstName: input.FirstName,
		LastName: input.LastName, Email: input.Email, PhoneNumber: input.PhoneNumber, Gender: input.Gender,
		ProfilePic: input.ProfilePic, Role_ID: 2}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	db := c.MustGet("db").(*gorm.DB)
	creatingUser := db.Create(&user)
	if creatingUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": creatingUser.Error})
		return
	}
	var userLimited ReturnSimpleUser
	userLimited.Username = user.Username
	userLimited.FirstName = user.FirstName
	userLimited.LastName = user.LastName
	userLimited.Email = user.Email

	c.JSON(http.StatusOK, gin.H{"data": userLimited})
}

// GET /users/:user_id
// Find a user
func FindUser(c *gin.Context) { // Get model if exist
	var user models.Commerce_User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /details
// Get Self User Detail
func SelfDetailUser(c *gin.Context) { // Get model if exist
	var user models.Commerce_User
	user_id, err := token.ExtractTokenID(c)
	db := c.MustGet("db").(*gorm.DB)
	if err = db.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /update
// Update self user
func UpdateSelfUser(c *gin.Context) {
	user_id, err := token.ExtractTokenID(c)
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.Commerce_User
	if err := db.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_User
	updatedInput.ProfilePic = input.ProfilePic
	updatedInput.Gender = input.Gender
	updatedInput.PhoneNumber = input.PhoneNumber
	updatedInput.Email = input.Email
	updatedInput.LastName = input.LastName
	updatedInput.FirstName = input.FirstName
	updatedInput.Password = input.Password
	updatedInput.Username = input.Username
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedInput.Password = string(hashedPassword)

	//remove spaces in username
	updatedInput.Username = html.EscapeString(strings.TrimSpace(input.Username))

	db.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH /users/:id
// Update a user
func UpdateUser(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.Commerce_User
	if err := db.Where("user_id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Commerce_User
	updatedInput.ProfilePic = input.ProfilePic
	updatedInput.Gender = input.Gender
	updatedInput.PhoneNumber = input.PhoneNumber
	updatedInput.Email = input.Email
	updatedInput.LastName = input.LastName
	updatedInput.FirstName = input.FirstName
	updatedInput.Password = input.Password
	updatedInput.Username = input.Username
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedInput.Password = string(hashedPassword)

	//remove spaces in username
	updatedInput.Username = html.EscapeString(strings.TrimSpace(input.Username))

	db.Model(&user).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var user models.Commerce_User
	if err := db.Where("user_id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
