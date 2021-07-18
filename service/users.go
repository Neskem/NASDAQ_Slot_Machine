package service

import (
	"NASDAQ_Slot_Machine/models"
	"NASDAQ_Slot_Machine/orm"
)

func GetOneUser(id int64) (*models.Users, error) {
	return orm.SelectOneUsers(id)
}
