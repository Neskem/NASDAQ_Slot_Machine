package controller

import (
	"NASDAQ_Slot_Machine/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UsersController struct {}

func NewUsersController() UsersController {
	return UsersController{}
}

func (u UsersController) GetOne(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id: ", id)

	userId, err := strconv.ParseInt(id, 10, 64);
	if (err != nil) {
		c.AbortWithStatus(400)
		fmt.Println(err.Error())
	}

	userOne, err := service.GetOneUser(userId)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, &userOne)
	}
	return
}
