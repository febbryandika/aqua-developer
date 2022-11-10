package config

import (
	"e-commerce/entity"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Database() {
	DB, err = gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&entity.Product{}, &entity.Cart{}, &entity.Checkout{}, &entity.CheckoutProduct{}, &entity.User{})
}

//Logger: logger.Default.LogMode(logger.Silent)
