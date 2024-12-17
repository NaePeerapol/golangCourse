package main

import (
	"fmt"
	"golangCourse/database"
	m "golangCourse/models"
	"golangCourse/routes"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_course",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected successfully!!")
	database.DBConn.AutoMigrate(&m.Dogs{})
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.WebAppRoutes(app)
	app.Listen(":3000")
}
