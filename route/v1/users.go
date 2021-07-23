package v1

import (
	"NASDAQ_Slot_Machine/controller"
	_ "NASDAQ_Slot_Machine/docs"
	"github.com/gin-gonic/gin"
)


func RouteUsers(r *gin.Engine) {
	posts := r.Group("/users")
	{
		posts.GET("/:id", controller.NewUsersController().GetOne)
		posts.POST("/register/", controller.NewUsersController().RegisterOne)
		posts.POST("/login/", controller.OldUsersController().LoginOne)
	}
}
