package service

import (
	"NASDAQ_Slot_Machine/dao"
	"NASDAQ_Slot_Machine/models"
)

func GetOneUser(id int64) (*models.Users, error) {
	return dao.SelectOneUsers(id)
}

func LoginUser(account string, password string) (*models.Users, error) {
	return dao.LoginOneUser(account, password)
}

func RegisterUser(account string, password string, email string) error {
	err := dao.RegisterOneUser(account, password, email)
	return err
}
