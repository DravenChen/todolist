package utils

import (
	"log"

	"github.com/DravenChen/todolist/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB *gorm.DB
var DB *gorm.DB

// InitDB init database connection
func InitDB() {
	var err error
	DB, err = gorm.Open( mysql.Open(config.C.DBConfig.ToDsn()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("database connected")
}