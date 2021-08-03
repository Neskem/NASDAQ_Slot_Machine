package main

import (
	"NASDAQ_Slot_Machine/database"
	_ "NASDAQ_Slot_Machine/docs"
	"NASDAQ_Slot_Machine/middleware"
	v1 "NASDAQ_Slot_Machine/route/v1"
	"fmt"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
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
	app.Use(middleware.CORSMiddleware())
	db, err1 := database.InitDb(dbConfig)
	if err1 != nil {
		fmt.Println("get db failed:", err)
		return
	}
	app.Use(database.Inject(db))
	app.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "hello " + name,
		})
	})
	// memoryStore := persist.NewMemoryStore(1 * time.Minute)
	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "redis:6379",
		DB: 0,
	}))
	v1.RouteUsers(app, redisStore)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	err2 := app.Run(":" + port)
	if err2 != nil {
		panic(err2)
	}
}



