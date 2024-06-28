package controllers

import (
	"ai_marketplace_go/models"
	"ai_marketplace_go/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var jwtKey = []byte("Diyarbek") // Replace with your own secret key

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Login godoc
// @Summary Login
// @Description Login
// @param user body models.LoginUser false "User object"
// @Produce application/json
// @tags users
// @Success 200 {string} json "Token"
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.User
	var loginData models.User
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &utils.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token, err := utils.GenerateJWT(claims, jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Profile(c *gin.Context) {
	username := c.MustGet("username").(string)
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"username":   user.Username,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"balance":    user.Balance,
		"email":      user.Email,
	})
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user
// @param user body models.SwaggerUser false "User object"
// @Produce application/json
// @tags users
// @Success 201 {string} string "User created successfully"
// @Router /register [post]
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	data := map[string]interface{}{
		"username":   user.Username,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"balance":    user.Balance,
		"email":      user.Email,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User created successful",
		"data":    data,
	})
}
