package database

import (
	"log"
	"os"

	"github.com/wayanputri/blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBConn *gorm.DB

func ConnectDB() {

	user := os.Getenv("db_user")
	pass := os.Getenv("db_password")
	host := os.Getenv("db_host")
	port := os.Getenv("db_port")
	name := os.Getenv("db_name")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		panic("Database connection failed.")
	}

	log.Println("Connection successfully")
	db.AutoMigrate(new(model.Blog))
	DBConn = db
}
