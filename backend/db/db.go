package db

import (
	"fmt"
	"log"
	"thomas-barth/SharedShopping/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	username := "thomas_barth"
	password := "TB-bt1a@"
	host := "127.0.0.1" // or "localhost"
	port := "3306"      // MySQL default port
	dbname := "SharedShoppingDB"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.ShoppingList{},
		&models.Item{},
		&models.CustomItem{},
	)

	if err != nil {
		log.Fatal(err)
	}

	DB = db

}
