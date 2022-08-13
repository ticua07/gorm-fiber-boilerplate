package book

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/ticua07/go-fiber-api/database"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {

	db := database.DBConn
	book := new(Book)
	err := c.BodyParser(book)

	log.Print(book.Title)
	log.Print(book.Author)
	log.Print(book.Rating)

	// If any of the properties are missing
	if book.Title == "" || book.Author == "" || book.Rating == 0 {
		return c.Status(422).SendString("Missing properties")
	}

	if err != nil { 
		return c.Status(500).SendString(err.Error())
	}

	db.Create(&book)
	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.First(&book, id)

	if book.Title == "" {
		c.Status(404).SendString("Book with id " + id + " not found")
	}

	db.Delete(&book)

	return c.SendString("Deleted book " + id)
}