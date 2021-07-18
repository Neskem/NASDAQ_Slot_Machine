package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)
func Initialize(dbConfig string) (*gorm.DB, error) {
	// dbConfig := "host=postgresql user=nasdaq password=password dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	Db, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})
	return Db, err
}
