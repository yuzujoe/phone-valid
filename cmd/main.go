package main

import (
	"fmt"
	"log"
	"os"
	"phone-valid/middleware"
	"phone-valid/mysql"
	"phone-valid/route"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Printf("Server started")

	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatal(err)
	}
	database := os.Getenv("MYSQL_DATABASE")

	path := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database,
	)

	db := mysql.Init(path)
	db.LogMode(true)
	mysql.Migrate(db)
	defer db.Close()

	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.Use(middleware.JwtAuth)

	route.Route(r)

	log.Fatal(r.Run())
}
