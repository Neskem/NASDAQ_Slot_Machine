package main

import (
	"NASDAQ_Slot_Machine/database"
	"NASDAQ_Slot_Machine/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)


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
	err2 := app.Run(":" + port)
	if err2 != nil {
		panic(err2)
	}
}



