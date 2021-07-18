package orm

import (
	"NASDAQ_Slot_Machine/database"
	"NASDAQ_Slot_Machine/models"
)


func SelectOneUsers(id int64) (*models.Users, error) {
	fields := []string{"id", "account", "email"}
	userOne:=&models.Users{}
	err := database.Db.Select(fields).Where("id=?", id).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}
