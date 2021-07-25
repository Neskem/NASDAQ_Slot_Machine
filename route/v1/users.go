package v1

import (
	"NASDAQ_Slot_Machine/controller"
	_ "NASDAQ_Slot_Machine/docs"
	"NASDAQ_Slot_Machine/middleware"
	"github.com/gin-gonic/gin"
)


func RouteUsers(r *gin.Engine) {
	posts := r.Group("/users")
	{
		posts.POST("/register/", controller.NewUsersController().RegisterOne)
		posts.POST("/login/", controller.OldUsersController().LoginOne)
	}

	auth := r.Group("/users/auth")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/:id", controller.NewUsersController().GetOne)
	}
}
