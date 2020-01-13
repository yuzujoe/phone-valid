package mysql

import (
	"fmt"
	"phone-valid/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init(path string) *gorm.DB {

	db, err := gorm.Open("mysql", path)
	if err != nil {
		panic(err)
	}

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success connected")

	DB = db

	return DB
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.AuthenticationCode{}).AddForeignKey("phone_number", "users(phone_number)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&models.UserProfile{}).AddForeignKey("user_id", "users(user_id)", "RESTRICT", "RESTRICT")
}
