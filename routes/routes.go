package routes

import (
	"ai_marketplace_go/controllers"
	"ai_marketplace_go/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // Swagger API documentation

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.Authenticate)
	{
		authorized.GET("/profile", controllers.Profile)
		authorized.POST("/upload_video", controllers.UploadVideo)
		authorized.POST("/upload_image", controllers.UploadImage)
		authorized.POST("/set_video_task", controllers.SetVideoTask)
		authorized.POST("/get_task_result", controllers.GetTaskResult)
		authorized.POST("/ai_models", controllers.CreateAiModel)
		authorized.GET("/ai_models", controllers.GetAiModels)
		authorized.GET("/ai_models/:id", controllers.GetAiModel)
		authorized.PUT("/ai_models/:id", controllers.UpdateAiModel)
		authorized.DELETE("/ai_models/:id", controllers.DeleteAiModel)

		authorized.POST("/ai_models/use", controllers.UseModel)

		authorized.POST("/contact", controllers.SendMessage)
		authorized.GET("/contact", controllers.GetMessages)
		authorized.GET("/contact/:id", controllers.GetMessage)
	}
}
