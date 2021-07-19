package v1

import (
	"NASDAQ_Slot_Machine/controller"
	"github.com/gin-gonic/gin"
)

func RouteUsers(r *gin.RouterGroup) {
	posts := r.Group("/ypa")
	{
		posts.GET("/users/:id", controller.NewUsersController().GetOne)
		posts.POST("/users/register/", controller.NewUsersController().RegisterOne)
	}
}
