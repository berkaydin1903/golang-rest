package middleware

import (
	"main/database"
	"main/models"
	"main/util"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(c *fiber.Ctx) error {
	c.Status(401)
	Authorization := c.Get("Authorization")

	id, _ := util.ParseJwt(Authorization)

	var user models.User

	database.DB.Where("id=? AND token=?", id, Authorization).First(&user)

	if user.Token == "" {
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "unauthenticated",
		})
	}

	return c.Next()
}
