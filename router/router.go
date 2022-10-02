package router

import (
	"github.com/gin-gonic/gin"
)

func start_app() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.user_register)
		userRouter.POST("/login", controllers.user_login)
		userRouter.PUT("/:userId", controllers.user_update)
		userRouter.DELETE("/:userId", controllers.user_delete)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.auth())

		photoRouter.POST("/", controllers.create_photo)
		photoRouter.GET("/", controllers.list_photo)
		photoRouter.PUT("/:photoId", middlewares.photo_auth(), controllers.update_photo)
		photoRouter.DELETE("/:photoId", middlewares.photo_auth(), controllers.update_photo)
	}

	return r
}
