package controllers

import (
	"main/database"
	"main/dto"
	"main/models"
	"strconv"

	"main/util"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string
	var UserRegisterDto dto.UserRegisterDto

	if err := c.BodyParser(&UserRegisterDto); err != nil {
		return err
	}

	if error := validate.Struct(&UserRegisterDto); error != nil {
		print(error)
		resp := util.ToErrResponse(error)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": resp.Errors,
		})

	}

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "GIRILEN SIFRELER ES DEGIL",
		})
	}
	user := models.User{
		UserName:     data["username"],
		Email:        data["email"],
		UserContacts: []models.UserContact{},
	}
	user.SetPassword(data["password"])
	database.DB.Create(&user)

	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)), user.Id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	Usertoken := models.User{
		Id:       uint(user.Id),
		UserName: user.UserName,
		Token:    token,
	}

	if err := c.BodyParser(&Usertoken); err != nil {
		return err
	}

	database.DB.Model(&Usertoken).Updates(Usertoken)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   Usertoken,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(404)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "KULLANICI BULUNAMADI",
		})

	}
	if err := user.ComparePassword(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "HATALI SIFRE",
		})
	}
	token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)), user.Id)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	Usertoken := models.User{
		Id:       uint(user.Id),
		UserName: user.UserName,
		Token:    token,
	}

	if err := c.BodyParser(&Usertoken); err != nil {
		return err
	}

	database.DB.Model(&Usertoken).Updates(Usertoken)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   Usertoken,
	})
}

type Claims struct {
	jwt.StandardClaims
}

func User(c *fiber.Ctx) error {

	Authorization := c.Get("Authorization")

	id, err := util.ParseJwt(Authorization)

	if err != nil {

		return c.JSON(fiber.Map{
			"status":  false,
			"message": "token doğrulanmadı",
		})
	}

	var user models.User

	database.DB.Where("id=?", id).First(&user)
	return c.JSON(fiber.Map{
		"status": true,
		"data":   user,
	})
}
func Logout(c *fiber.Ctx) error {

	Authorization := c.Get("Authorization")

	id, err := util.ParseJwt(Authorization)
	if err != nil {

		return c.JSON(fiber.Map{
			"status":  false,
			"message": "token doğrulanmadı",
		})
	}

	database.DB.Model(&models.User{}).Where("id=? AND token=?", id, Authorization).Update("token", nil)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   nil,
	})
}
