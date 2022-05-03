package controllers

import (
	"main/database"
	"main/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetUserContactById(c *fiber.Ctx) error {

	// id, _ := strconv.Atoi(c.Params("id"))

	UserContacts := []models.UserContact{}
	database.DB.Find(&UserContacts)

	return c.JSON(&UserContacts)

}
func AddUserContact(c *fiber.Ctx) error {
	var usercontact models.UserContact

	if err := c.BodyParser(&usercontact); err != nil {
		return err
	}
	database.DB.Create(&usercontact)
	return c.JSON(fiber.Map{
		"status": true,
	})
}
func GetUserContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	UserContact := models.UserContact{
		Id: uint(id),
	}
	database.DB.Find(&UserContact)
	return c.JSON(&UserContact)
}
func UpdateUserContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	UserContact := models.UserContact{
		Id: uint(id),
	}

	if err := c.BodyParser(&UserContact); err != nil {
		return err
	}

	database.DB.Model(&UserContact).Updates(UserContact)

	return c.JSON(UserContact)
}
func DeleteUserContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	UserContact := models.UserContact{
		Id: uint(id),
	}

	database.DB.Delete(&UserContact)

	return nil
}
