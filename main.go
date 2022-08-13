package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/ticua07/go-fiber-api/book"
	"github.com/ticua07/go-fiber-api/database"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/books", book.GetBooks)
	app.Get("/book/:id", book.GetBook)
	app.Post("/addBook", book.NewBook)
	app.Delete("/deleteBook/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Database connection successfully initialized")

	database.DBConn.AutoMigrate(&book.Book{}) // Adapt sql table to Book struct
	log.Print("Database migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close() // close on exit or error

	app.Use(func(c *fiber.Ctx) error {
		// log requests
		log.Print(c.Hostname() + " accessed " + c.Path())
		return c.Next()
	})

	SetupRoutes(app)

	log.Fatal(app.Listen("localhost:3000"))
}