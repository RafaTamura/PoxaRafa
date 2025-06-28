package routes

import (
	"poxa_rafa/controllers"
	"poxa_rafa/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "poxa_rafa/docs"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	// rota pública
	router.POST("/login", controllers.Login)
	router.GET("/posts", controllers.GetPosts)
	router.GET("/posts/:id", controllers.GetPostByID)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// rotas protegidas
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/posts", controllers.CreatePost)
		auth.PUT("/posts/:id", controllers.UpdatePost)
		auth.DELETE("/posts/:id", controllers.DeletePost)
	}

	return router
}
