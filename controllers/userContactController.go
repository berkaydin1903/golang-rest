package controllers

import (
	"main/database"
	"main/dto"
	"main/models"
	"main/util"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func GetUserContactById(c *fiber.Ctx) error {

	id, _ := strconv.Atoi(c.Params("id"))
	UserContact := models.UserContact{
		Id: uint(id),
	}
	database.DB.Find(&UserContact)
	return c.JSON(&UserContact)

}

var validate = validator.New()

func AddUserContact(c *fiber.Ctx) error {
	var usercontactdto dto.UserContactDto
	Authorization := c.Get("Authorization")
	id, _ := util.ParseJwt(Authorization)

	if err := c.BodyParser(&usercontactdto); err != nil {
		return err
	}

	if error := validate.Struct(&usercontactdto); error != nil {
		resp := util.ToErrResponse(error)
		return c.JSON(fiber.Map{
			"status": false,
			"error":  resp.Errors,
		})
	}
	usercontact := models.UserContact{
		Name:   usercontactdto.Name,
		Phone:  usercontactdto.Phone,
		UserId: uint(id),
	}
	database.DB.Create(&usercontact)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   usercontact,
	})
}
func GetUserContact(c *fiber.Ctx) error {
	Authorization := c.Get("Authorization")
	id, _ := util.ParseJwt(Authorization)
	var UserContact []models.UserContact
	database.DB.Where("user_id=?", id).Find(&UserContact)
	c.Status(200)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   UserContact,
	})
}
func UpdateUserContact(c *fiber.Ctx) error {
	var usercontactdto dto.UserContactDto
	Authorization := c.Get("Authorization")
	user_id, _ := util.ParseJwt(Authorization)

	if err := c.BodyParser(&usercontactdto); err != nil {
		return err
	}

	if error := validate.Struct(&usercontactdto); error != nil {
		resp := util.ToErrResponse(error)
		return c.JSON(fiber.Map{
			"status": false,
			"error":  resp.Errors,
		})
	}
	id, _ := strconv.Atoi(c.Params("id"))

	UserContact := models.UserContact{
		Id:     uint(id),
		Name:   usercontactdto.Name,
		Phone:  usercontactdto.Phone,
		UserId: user_id,
	}

	database.DB.Model(&UserContact).Updates(UserContact)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   UserContact,
	})
}
func DeleteUserContact(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	Authorization := c.Get("Authorization")
	user_id, _ := util.ParseJwt(Authorization)

	UserContact := models.UserContact{
		Id:     uint(id),
		UserId: user_id,
	}

	database.DB.Delete(&UserContact)
	c.Status(200)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   nil,
	})
}
