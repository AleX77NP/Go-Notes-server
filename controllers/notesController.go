package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aleksandarmilanovic/fb-training/db"
	"github.com/aleksandarmilanovic/fb-training/models"
    ) 

//GetNotes ...
func GetNotes(c *fiber.Ctx) error {
	db := db.DBConn

	var notes []models.Note

	db.Find(&notes)
	return c.JSON(notes)
}

//CreateNote ...
func CreateNote(c *fiber.Ctx) error {
	db := db.DBConn

	note := new(models.Note)

	if err := c.BodyParser(note); err != nil {
		return c.Status(400).SendString("Invalid data")
	}

	if note.Text == "" || note.Title == "" || note.User == "" {
		return c.Status(400).SendString("Invalid data")
	}

	db.Create(&note)
	return c.JSON(note)
}

//DeleteNote ...
func DeleteNote(c *fiber.Ctx) error {
	db := db.DBConn

	id := c.Params("id")
	var note models.Note
	db.First(&note, id)

	if note.Title == "" {
		return c.Status(404).SendString("Note not found")
	}
	db.Delete(&note)
	return c.Status(200).SendString("Note deleted.")

}

//UpdateNote ...
func UpdateNote(c *fiber.Ctx) error {
	db := db.DBConn

	id := c.Params("id")
	noteChange := new(models.NoteChange)
	var note models.Note
	
	if err := c.BodyParser(noteChange); err != nil {
		return c.Status(400).SendString("Invalid info")
	}
	db.First(&note, id)
	note.Title = noteChange.Title
	note.Text = noteChange.Text
	db.Save(&note)

	return c.Status(200).SendString("Note updated.")

}