package app

import (
	"fmt"
	"task/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_Username string
	DB_Password string
	DB_Host     string
	DB_Port     string
	DB_Database string
}

func NewDB(config *ConfigDB) *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.DB_Host,
		config.DB_Username,
		config.DB_Password,
		config.DB_Database,
		config.DB_Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicHandling(err)
	fmt.Printf("database Successfully connected!")
	return db
}
