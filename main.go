package main

import (
	"ai_marketplace_go/config"
	_ "ai_marketplace_go/docs"
	"ai_marketplace_go/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// @title Go Gin Auth API
// @version 1.0
// @description This is a sample server Go Gin Auth server.

// @host marketplace.araltech.tech
// @BasePath /
func main() {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db := config.InitDB()

	routes.SetupRoutes(r, db)

	r.Run(":8089")
}
