package main

import (
	"fmt"
	"log"

	"github.com/aleksandarmilanovic/fb-training/db"
	"github.com/aleksandarmilanovic/fb-training/controllers"
	"github.com/aleksandarmilanovic/fb-training/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDatabase() {
	var err error
	db.DBConn, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("DB Connection error!")
	}
	fmt.Println("DB Connection working...")
	db.DBConn.AutoMigrate(&models.User{})
	db.DBConn.AutoMigrate(&models.Note{})
}

func setUpRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Server works...")
	})
	app.Get("/api/users", controllers.GetUsers)
	app.Post("/api/users/signup", controllers.Signup)
	app.Post("/api/users/signin", controllers.Signin)

	app.Get("/api/users/notes/:username", controllers.GetUserNotes)

	app.Get("/api/notes", controllers.GetNotes)
	app.Post("/api/notes", controllers.CreateNote)
	app.Delete("/api/notes/:id", controllers.DeleteNote)
	app.Put("api/notes/:id", controllers.UpdateNote)
}


func main() {

	app := fiber.New()
	initDatabase()

	setUpRoutes(app)

	log.Fatal(app.Listen(":8000"))
}