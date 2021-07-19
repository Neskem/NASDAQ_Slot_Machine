package orm

import (
	"NASDAQ_Slot_Machine/database"
	"NASDAQ_Slot_Machine/models"
	"fmt"
)


var UserFields = []string{"id", "account", "email"}

func SelectOneUsers(id int64) (*models.Users, error) {
	userOne:=&models.Users{}
	err := database.Db.Select(UserFields).Where("id=?", id).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}

func LoginOneUser(account string, password string) (*models.Users, error) {
	userOne:=&models.Users{}
	err := database.Db.Select(UserFields).Where("account=?", account).Where("password=?", password).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}

func RegisterOneUser(account string, password string, email string) error {
	if CheckOneUser(account) {
		return fmt.Errorf("User exists.")
	}
	user := models.Users{
		Account: account,
		Password: password,
		Email: email,
	}
	insertErr := database.Db.Model(&models.Users{}).Create(&user).Error
	return insertErr
}

func CheckOneUser(account string) bool {
	result := false
	var user models.Users

	dbResult := database.Db.Where("account = ?", account).Find(&user)
	if dbResult.Error != nil {
		fmt.Printf("Get User Info Failed:%v\n", dbResult.Error)
	} else {
		result = true
	}
	return result

}