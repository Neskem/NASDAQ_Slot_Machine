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

func CreateUsersController() UsersController {
	return UsersController{}
}

// GetOne RouteUsers @Summary
// @Tags users
// @version 1.0
// @produce text/plain
// @param id path int true "id" default(1)
// @Success 200 string string successful return data
// @Router /users/{id} [get]
func (u UsersController) GetOne(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println("id: ", id)

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
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

type Login struct {
	Account string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Account string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func (u UsersController) LoginOne(c *gin.Context) {
	form := &Login{}
	if c.Bind(form) == nil {
		fmt.Println(form.Account, form.Password)
	}

	userOne, err := service.LoginUser(form.Account, form.Password)
	if err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err.Error())
	} else {
		c.JSON(http.StatusOK, &userOne)
	}
	return
}

// RegisterOne GetOne RouteUsers @Summary
// @Tags users
// @version 1.0
// @produce application/json
// @param body body Register true "JSON data" default({"account": "111", "password": "222", "email": "333"})
// @Success 200 string string successful return value
// @Router /users/register [post]
func(u UsersController) RegisterOne(c *gin.Context) {
	var form Register
	bindErr := c.BindJSON(&form)
	if bindErr == nil {
		// User regist
		err := service.RegisterUser(form.Account, form.Password, form.Email)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "success Register",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "Register Failed" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "Failed to parse register data" + bindErr.Error(),
			"data":   nil,
		})
	}
}