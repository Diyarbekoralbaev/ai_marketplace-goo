package main

import (
	"ai_marketplace_go/config"
	_ "ai_marketplace_go/docs"
	"ai_marketplace_go/routes"
	"github.com/gin-gonic/gin"
)

// @title Go Gin Auth API
// @version 1.0
// @description This is a sample server Go Gin Auth server.

// @host localhost:8089
// @BasePath /
func main() {
	r := gin.Default()
	db := config.InitDB()

	routes.SetupRoutes(r, db)

	r.Run(":8089")
}
