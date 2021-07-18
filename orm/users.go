package gorm

import (
	"NASDAQ_Slot_Machine/models"
	"gorm.io/gorm"
)


func batchInsertUsers(db *gorm.DB, users []*models.Users) {
	db.CreateInBatches(users, 100)
}