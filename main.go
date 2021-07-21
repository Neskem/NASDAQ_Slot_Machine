package main

import (
	v1 "NASDAQ_Slot_Machine/api/v1"
	"NASDAQ_Slot_Machine/database"
	_ "NASDAQ_Slot_Machine/docs"
	"NASDAQ_Slot_Machine/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
)


// @title Gin swagger
// @version 1.0
// @description Gin swagger

// @contact.name Flynn Sun

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	dbConfig := os.Getenv("DB_CONFIG")
	app := gin.Default()
	db, err1 := database.Initialize(dbConfig)
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	db.Debug().AutoMigrate(&models.Users{})
	migrator := db.Migrator()
	has := migrator.HasTable(&models.Users{})
	if !has {
		fmt.Println("table not exist")
	} else {
		fmt.Println("table already exist")
	}
	app.Use(database.Inject(db))
	app.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello " + name,
		})
	})
	v1.RouteUsers(app)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	err2 := app.Run(":" + port)
	if err2 != nil {
		panic(err2)
	}
}



