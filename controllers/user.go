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
	Username string `json:"username" binding:"required" example:"username123"`
	Password string `json:"password" binding:"required" example:"password123!"`
} // @name LoginUserInput

type CreateUserInput struct {
	Username    string `json:"username" binding:"required" example:"username123"`
	Password    string `json:"password" binding:"required" example:"password123!"`
	FirstName   string `json:"first_name" binding:"required" example:"John"`
	LastName    string `json:"last_name" binding:"required" example:"Asep"`
	Email       string `json:"email" example:"johnasep@mail.com"`
	PhoneNumber string `json:"phone_number" example:"08812314555"`
	Gender      string `json:"gender" example:"L"`
	ProfilePic  string `json:"profile_pic" example:"https://usr.co.id/img.jpg"`
} // @name CreateUserInput

type UpdateUserInput struct {
	Username    string `json:"username" binding:"required" example:"username123"`
	Password    string `json:"password" binding:"required" example:"password123!"`
	FirstName   string `json:"first_name" binding:"required" example:"John"`
	LastName    string `json:"last_name" binding:"required" example:"Asep"`
	Email       string `json:"email" example:"johnasep@mail.com"`
	PhoneNumber string `json:"phone_number" example:"08812314555"`
	Gender      string `json:"gender" example:"L"`
	ProfilePic  string `json:"profile_pic" example:"https://usr.co.id/img.jpg"`
} // @name UpdateUserInput

type ReturnSimpleUser struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
} // @name ReturnSimpleUser

// LoginUser godoc
// @Summary Login as as user.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags user
// @Param Body body LoginUserInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/login [post]
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

// FindUsers godoc
// @Summary Get all user.
// @Description Get a list of user registered in the system.
// @Tags admin
// @Produce json
// @Security BearerToken
// @param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]interface{}
// @Router /api/admin/users [get]
func FindUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.Commerce_User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// CreateUser godoc
// @Summary Create/Register a user.
// @Description Creating/registering a user from public access.
// @Tags user
// @Param Body body CreateUserInput true "the body to create a user"
// @Produce json
// @Success 200 {object} ReturnSimpleUser
// @Router /api/register [post]
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

// FindUser godoc
// @Summary Get one user.
// @Description Get a user by its id.
// @Tags admin
// @Produce json
// @Security BearerToken
// @Param user_id path string true "The user id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]interface{}
// @Router /api/admin/users/:user_id [get]
func FindUser(c *gin.Context) { // Get model if exist
	var user models.Commerce_User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("user_id = ?", c.Param("user_id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// SelfDetailUser godoc
// @Summary Get detail user.
// @Description Get a user its own detail.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]interface{}
// @Router /api/user/details [get]
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

// UpdateSelfUser godoc
// @Summary Update user data.
// @Description Update user its own data.
// @Tags user
// @Produce json
// @Security BearerToken
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]interface{}
// @Router /api/user/update [get]
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

// UpdateUser godoc
// @Summary Update one user.
// @Description Update a user by its id.
// @Tags admin
// @Produce json
// @Security BearerToken
// @Param user_id path string true "The user id"
// @Param Body body UpdateUserInput true "the body to update a user"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]interface{}
// @Router /api/admin/users/:user_id [patch]
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

// DeleteUser godoc
// @Summary Delete one user.
// @Description Delete a user by its id.
// @Tags admin
// @Produce json
// @Security BearerToken
// @Param user_id path string true "The user id"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Success 200 {object} map[string]interface{}
// @Router /api/admin/users/:user_id [delete]
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
