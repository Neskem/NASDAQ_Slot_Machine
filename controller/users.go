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

func(u UsersController) RegisterOne(c *gin.Context) {
	var form Register
	bindErr := c.BindJSON(&form)
	if bindErr == nil {
		// 用户注册
		err := service.RegisterUser(form.Account, form.Password, form.Email)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "success ",
				"data":   nil,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "注册失败" + err.Error(),
				"data":   nil,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "用户注册解析数据失败" + bindErr.Error(),
			"data":   nil,
		})
	}
}