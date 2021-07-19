package service

import (
	"NASDAQ_Slot_Machine/models"
	"NASDAQ_Slot_Machine/orm"
)

func GetOneUser(id int64) (*models.Users, error) {
	return orm.SelectOneUsers(id)
}

func LoginUser(account string, password string) (*models.Users, error) {
	return orm.LoginOneUser(account, password)
}

func RegisterUser(account string, password string, email string) error {
	err := orm.RegisterOneUser(account, password, email)
	return err
}
