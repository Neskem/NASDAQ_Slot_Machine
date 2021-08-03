package controller

import (
	"NASDAQ_Slot_Machine/middleware"
	"NASDAQ_Slot_Machine/models"
	"NASDAQ_Slot_Machine/service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type UsersController struct {}

func NewUsersController() UsersController {
	return UsersController{}
}

func OldUsersController() UsersController {
	return UsersController{}
}

func CreateUsersController() UsersController {
	return UsersController{}
}

// GetOne RouteUsers @Summary
// @Tags users
// @version 1.0
// @produce application/json
// @param token header string true "token"
// @param id path int true "id" default(1)
// @Success 200 string string successful return data
// @Router /users/auth/{id} [get]
func (u UsersController) GetOne(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	if claims == nil {
		c.AbortWithStatus(401)
	}
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
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token is effective",
			"data":   claims,
			"user": &userOne,
		})
	}
	return
}

type Login struct {
	Account string `json:"account" example:"000" binding:"required"`
	Password string `json:"password" example:"000" binding:"required"`
}

type Register struct {
	Account string `json:"account" binding:"required" example:"account"`
	Password string `json:"password" binding:"required" example:"password"`
	Email string `json:"email" binding:"required" example:"test123@gmail.com"`
}

type LoginResult struct {
	Account string `json:"account" binding:"required"`
	Token string `json:"token" binding:"required"`
}

type LoginResponse struct {
    Status int64 `json:"status" example:"-1"`
    Msg string `json:"msg" example:"Successfully login."`
    Data LoginResult `json:"data"`
}

type RegisterResponse struct {
	Status int64 `json:"status" example:"0"`
	Msg string `json:"msg" example:"Successfully login."`
	Data string `json:"data"`
}
// LoginOne RouteUsers @Summary
// @Tags users
// @version 1.0
// @produce application/json
// @param body body Login true "JSON data"
// @Success 200 string string successful login
// @Router /users/login/ [post]
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
		generateToken(c, userOne)
	}
	return
}

// RegisterOne RouteUsers @Summary
// @Tags users
// @version 1.0
// @produce application/json
// @param register body Register true "register"
// @Success 200 {array} RegisterResponse
// @Router /users/register/ [post]
func(u UsersController) RegisterOne(c *gin.Context) {
	var form Register
	bindErr := c.BindJSON(&form)
	if bindErr == nil {
		// User regist
		err := service.RegisterUser(form.Account, form.Password, form.Email)

		if err == nil {
			c.JSON(http.StatusOK, RegisterResponse {
				Status: 0,
				Msg: "Success Register",
				Data: "",
			})
		} else {
			c.JSON(http.StatusOK, RegisterResponse {
				Status: -1,
				Msg: "Register Failed" + err.Error(),
				Data: "",
			})
		}
	} else {
		c.JSON(http.StatusOK, RegisterResponse {
			Status: -1,
			Msg: "Failed to parse register data" + bindErr.Error(),
			Data: "",
		})
	}
}


// token生成器
// md 為上面定義好的middleware中介軟體
func generateToken(c *gin.Context, user *models.Users) {
	// 構造SignKey: 簽名和解簽名需要使用一個值
	j := middleware.NewJWT()

	// 構造使用者claims資訊(負荷)
	claims := middleware.CustomClaims{
		Account:  user.Account,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // Effective date
			ExpiresAt: int64(time.Now().Unix() + 3600), // Expired date
			Issuer:    "Flynn.Sun",                    // Signer
		},
	}

	// Generate token from claim
	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, LoginResponse{
			-1,
			err.Error(),
			 LoginResult{},
		})
	}

	log.Println("Token: ", token)
	data := LoginResult {
		Account:  user.Account,
		Token: token,
	}

	c.JSON(http.StatusOK, LoginResponse{
		0,
		"Successfully login.",
		data,
	})
}
