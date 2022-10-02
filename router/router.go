package router

import (
	"rakaminbtpn/controllers"
	"rakaminbtpn/middleware"

	"github.com/gin-gonic/gin"
)

func Start_app() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.User_register)
		userRouter.POST("/login", controllers.User_login)
		userRouter.PUT("/:userId", controllers.User_update)
		userRouter.DELETE("/:userId", controllers.User_delete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Auth())

		photoRouter.POST("/", controllers.Create_photo)
		photoRouter.GET("/", controllers.List_photo)
		photoRouter.PUT("/:photoId", middleware.Photo_auth(), controllers.Update_photo)
		photoRouter.DELETE("/:photoId", middleware.Photo_auth(), controllers.Update_photo)
	}

	return r
}
