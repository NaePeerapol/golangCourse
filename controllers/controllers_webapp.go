package controllers

import (
	"golangCourse/database"
	m "golangCourse/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HelloTest(c *fiber.Ctx) error {
	return c.SendString("Peerapol Chanviriyapreeda")
}

func HelloTestV2(c *fiber.Ctx) error {
	return c.SendString("Nae Peerapol V2")
}

func HelloTestV3(c *fiber.Ctx) error {
	return c.SendString("This is the first page")
}

func PersonTest(c *fiber.Ctx) error {
	p := new(m.Person)

	if err := c.BodyParser(p); err != nil {
		return err
	}

	log.Println(p.Firstname)
	log.Println(p.Lastname)
	str := p.Firstname + " " + p.Lastname
	return c.JSON(str)
}

func ParamsTest(c *fiber.Ctx) error {
	str := "This is the word that you write on the URL ==> " + c.Params("name")
	return c.JSON(str)
}

func CreateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func DeleteDog(c *fiber.Ctx) error {
	db := database.DBConn //เชื่อมต่อ database
	id := c.Params("id") //รัับค่าจาก parameter :id
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.Status(404).SendString("ไม่พบข้อมูลที่ผู้ใช้ต้องการ")
	}

	return c.Status(200).SendString("ลบข้อมูลเรียบร้อยแล้ว!!!")
}

func GetDeleteDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	result := db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.Status(200).JSON(dogs)
}

func GetAllDogs(c *fiber.Ctx) error {
	db := database.DBConn.Debug()
	var dogs []m.Dogs

	db.Find(&dogs) //SELECT * FROM `dogs` WHERE `dogs`.`deleted_at` IS NULL

	return c.Status(200).JSON(dogs)

}

func GetDogId(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Find(&dog, id)

	if result.RowsAffected == 0 {
		return c.Status(404).SendString("ไม่พบข้อมูลที่ต้องการ")
	}

	return c.Status(200).JSON(dog)
}