package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/controllers"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/database"
	"github.com/AgusRakhmatHaryanto/task5-pbi-btpns-AgusRakhmatHaryanto/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/uploads", "./uploads")

	db := database.GetDB()

	userController := controllers.NewUserController(db)
	photoController := controllers.NewPhotoController(db)

	api := router.Group("/api/v1")

	userRoutes := api.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.PUT("/update/:userId",  userController.Update)
		userRoutes.DELETE("/delete", userController.Delete)
	}

	photoRoutes := api.Group("/photos")
	{
		photoRoutes.GET("/", middlewares.AuthMiddleware(db), photoController.GetPhotos)
		photoRoutes.POST("/", middlewares.AuthMiddleware(db), photoController.CreatePhoto)
		photoRoutes.DELETE("/delete", middlewares.AuthMiddleware(db), photoController.DeletePhoto)
		photoRoutes.PUT("/update", middlewares.AuthMiddleware(db), photoController.UpdatePhoto)
	}
	return router
}