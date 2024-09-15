package database

import (
	"fmt"
	"log"
	giftModel "mit-api/pkg/gift/model"
	spinWheelModel "mit-api/pkg/spin-wheel/model"
	tourModel "mit-api/pkg/tour/model"
	userGiftModel "mit-api/pkg/user-gift/model"
	userModel "mit-api/pkg/user/model"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	DBInstance *gorm.DB
)

func New() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	DBInstance = db
	db.AutoMigrate(&userModel.User{}, &tourModel.Tour{}, &userGiftModel.UserGift{}, &spinWheelModel.SpinWheel{}, &giftModel.Gift{})
}
