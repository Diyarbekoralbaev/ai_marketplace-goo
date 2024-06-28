package controllers

import (
	"ai_marketplace_go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SendMessage godoc
// @Summary Contact
// @Description Contact
// @param contact body models.SwaggerContact true "Contact object"
// @Produce application/json
// @tags contact
// @Success 201 {object} models.SwaggerContact
// @Router /contact [post]
func SendMessage(c *gin.Context) {
	var contact models.Contact
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&contact).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error sending message"})
		return
	}

	c.JSON(201, contact)
}

// GetMessages godoc
// @Summary Get all messages
// @Description Get all messages
// @Produce application/json
// @tags contact
// @Success 200 {array} models.SwaggerContact
// @Router /contact [get]
func GetMessages(c *gin.Context) {
	var contacts []models.Contact
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&contacts)

	c.JSON(200, contacts)
}

// GetMessage godoc
// @Summary Get message
// @Description Get message
// @Produce application/json
// @tags contact
// @Param id path int true "Message ID"
// @Success 200 {object} models.SwaggerContact
// @Router /contact/{id} [get]
func GetMessage(c *gin.Context) {
	var contact models.Contact
	db := c.MustGet("db").(*gorm.DB)
	if err := db.First(&contact, c.Param("id")).Error; err != nil {
		c.JSON(404, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(200, contact)
}
