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
	r.POST("/ai_models", controllers.CreateAiModel)
	r.GET("/ai_models", controllers.GetAiModels)
	r.GET("/ai_models/:id", controllers.GetAiModel)
	r.PUT("/ai_models/:id", controllers.UpdateAiModel)
	r.DELETE("/ai_models/:id", controllers.DeleteAiModel)

	r.POST("/ai_models/use", controllers.UseModel)

	r.POST("/contact", controllers.SendMessage)
	r.GET("/contact", controllers.GetMessages)
	r.GET("/contact/:id", controllers.GetMessage)

	r.POST("/upload_video", controllers.UploadVideo)
	r.POST("/upload_image", controllers.UploadImage)
	r.POST("/set_video_task", controllers.SetVideoTask)
	r.POST("/get_task_result", controllers.GetTaskResult)

	authorized := r.Group("/")
	authorized.Use(middleware.Authenticate)
	{
		authorized.GET("/profile", controllers.Profile)
	}
}
