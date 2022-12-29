package db

import (
	"fmt"
	"github.com/AzizRahimov/e-wallet/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// StartDbConnection - connect to db
func StartDbConnection() *gorm.DB {
	settingParams := utils.AppSettings.PostgresParams

	connString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		settingParams.Server, settingParams.Port,
		settingParams.User, settingParams.DataBase,
		settingParams.Password)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		log.Fatal("Couldn't connect to database", err.Error())
	}

	return db
}
