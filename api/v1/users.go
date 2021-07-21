package v1

import (
	"NASDAQ_Slot_Machine/controller"
	_ "NASDAQ_Slot_Machine/docs"
	"github.com/gin-gonic/gin"
)

// RouteUsers @Summary 說Hello
// @Id 1
// @Tags Hello
// @version 1.0
// @produce text/plain
// @Success 200 string string 成功後返回的值
// @Router /hello [get]
func RouteUsers(r *gin.Engine) {
	posts := r.Group("/ypa")
	{
		posts.GET("/users/:id", controller.NewUsersController().GetOne)
		posts.POST("/users/register/", controller.NewUsersController().RegisterOne)
	}
}
