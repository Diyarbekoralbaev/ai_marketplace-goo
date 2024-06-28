package controllers

import (
	"ai_marketplace_go/marketplace_models"
	"ai_marketplace_go/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAiModels godoc
// @Summary Get all AI models
// @Description Get all AI models
// @Produce application/json
// @tags ai_models
// @Success 200 {array} models.SwaggerAiModel
// @Router /ai_models [get]
func GetAiModels(c *gin.Context) {
	var aiModels []models.AiModels
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&aiModels)

	c.JSON(200, aiModels)
}

// CreateAiModel godoc
// @Summary Create AI model
// @Description Create AI model
// @param aiModel body models.SwaggerAiModel true "AI model object"
// @Produce application/json
// @tags ai_models
// @Success 201 {object} models.SwaggerAiModel
// @Router /ai_models [post]
func CreateAiModel(c *gin.Context) {
	var aiModel models.AiModels
	if err := c.ShouldBindJSON(&aiModel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&aiModel).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error creating ai model"})
		return
	}

	c.JSON(201, aiModel)
}

// GetAiModel godoc
// @Summary Get AI model
// @Description Get AI model
// @Produce application/json
// @tags ai_models
// @Param id path int true "AI model ID"
// @Success 200 {object} models.SwaggerAiModel
// @Router /ai_models/{id} [get]
func GetAiModel(c *gin.Context) {
	var aiModel models.AiModels
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&aiModel).Error; err != nil {
		c.JSON(404, gin.H{"error": "Ai model not found"})
		return
	}

	c.JSON(200, aiModel)
}

// UpdateAiModel godoc
// @Summary Update AI model
// @Description Update AI model
// @param id path int true "AI model ID"
// @param aiModel body models.SwaggerAiModel true "AI model object"
// @Produce application/json
// @tags ai_models
// @Success 200 {object} models.SwaggerAiModel
// @Router /ai_models/{id} [put]
func UpdateAiModel(c *gin.Context) {
	var aiModel models.AiModels
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&aiModel).Error; err != nil {
		c.JSON(404, gin.H{"error": "Ai model not found"})
		return
	}

	if err := c.ShouldBindJSON(&aiModel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db.Save(&aiModel)
	c.JSON(200, aiModel)
}

// DeleteAiModel godoc
// @Summary Delete AI model
// @Description Delete AI model
// @Produce application/json
// @tags ai_models
// @Param id path int true "AI model ID"
// @Success 200 {string} string "AI model deleted successfully"
// @Router /ai_models/{id} [delete]
func DeleteAiModel(c *gin.Context) {
	var aiModel models.AiModels
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&aiModel).Error; err != nil {
		c.JSON(404, gin.H{"error": "Ai model not found"})
		return
	}

	db.Delete(&aiModel)
	c.JSON(200, gin.H{"message": "Ai model deleted successfully"})
}

// UseModel godoc
// @Summary Use AI model
// @Description Use AI model
// @Produce multipart/form-data
// @Param text formData string true "Text to analyze"
// @Param image formData file false "Image to analyze"
// @tags ai_models
func UseModel(c *gin.Context) {
	apiKey := "AIzaSyBHENaVP_KEfM7Bm0fuLAfxllJ8MGECpms"
	if c.PostForm("text") == "" {
		c.JSON(400, gin.H{"error": "Text is required"})
		return
	}
	if c.PostForm("image_path") == "" {
		text := c.PostForm("text")
		content := marketplace_models.GenerateText(apiKey, text)
		c.JSON(200, gin.H{"content": content})
	}
}
