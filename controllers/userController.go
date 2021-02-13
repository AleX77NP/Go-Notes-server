package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/aleksandarmilanovic/fb-training/db"
	"github.com/aleksandarmilanovic/fb-training/models"
    ) 


//GetUsers ...
func GetUsers(c *fiber.Ctx) error {
	db := db.DBConn

	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}

//Signup ...
func Signup(c *fiber.Ctx) error {
	db := db.DBConn

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(503).Send([]byte(err.Error()))
	}

	db.Create(&user)
	return c.JSON(user)
}

//Signin ...
func Signin(c *fiber.Ctx) error {
	db := db.DBConn

	user := new(models.User)
	userInfo := new(models.LoginInfo)
	if err := c.BodyParser(userInfo); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	if userInfo.Username == "" || userInfo.Password == "" {
		return c.Status(400).SendString("User info not correct")
	}

	db.Where("username = ?", userInfo.Username).First(&user)

	if user.Username == "" {
		return c.Status(404).SendString("Not found")
	}
	
	if userInfo.Password == user.Password {
		return c.SendStatus(200)
	}
	return c.Status(400).SendString("Wrong password")
}

//GetUserNotes ...
func GetUserNotes(c *fiber.Ctx) error {
	db := db.DBConn

	username := c.Params("username")
	var notes []models.Note

	db.Where("user = ?", username).Find(&notes)

	return c.JSON(notes)


}