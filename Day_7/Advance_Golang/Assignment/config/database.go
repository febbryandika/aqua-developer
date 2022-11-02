package config

import (
	"os"
	"user-management/entity"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		panic(err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&entity.User{})
}
