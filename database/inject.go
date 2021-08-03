package database

import (
	"NASDAQ_Slot_Machine/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Inject injects database to gin context
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func InitDb(config string) (*gorm.DB, error) {
	db, err1 := Initialize(config)
	if err1 != nil {
		return nil, err1
	}
	db.Debug().AutoMigrate(&models.Users{})
	migrator := db.Migrator()
	has := migrator.HasTable(&models.Users{})
	if !has {
		fmt.Println("table not exist")
	} else {
		fmt.Println("table already exist")
	}
	return db, err1
}