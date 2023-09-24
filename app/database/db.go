package database

import (
	"fmt"
	"referrer/app/config"
	"referrer/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	mysqlHost := config.Getenv("MYSQL_HOST")
	mysqlPort := config.Getenv("MYSQL_PORT")
	mysqlUser := config.Getenv("MYSQL_USER")
	mysqlPassword := config.Getenv("MYSQL_PASSWORD")
	mysqlDatabase := config.Getenv("MYSQL_DATABASE")

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabase)), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(models.User{}, models.Product{}, models.Link{}, models.Order{}, models.OrderItem{})
}
