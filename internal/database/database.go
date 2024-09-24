package database

import (
	"fmt"
	giftModel "mit-api/pkg/gift/model"
	rbacModel "mit-api/pkg/rbac/model"
	spinWheelModel "mit-api/pkg/spin-wheel/model"
	tourModel "mit-api/pkg/tour/model"
	userGiftModel "mit-api/pkg/user-gift/model"
	userModel "mit-api/pkg/user/model"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DBInstance *gorm.DB
)

func Connect(host string, username string, password string, database string, port string) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	DBInstance = db
	db.AutoMigrate(&userModel.User{}, &tourModel.Tour{}, &userGiftModel.UserGift{}, &spinWheelModel.SpinWheel{}, &giftModel.Gift{}, &rbacModel.Role{}, &rbacModel.Permission{}, &rbacModel.RolePermission{})

	return db, nil
}
