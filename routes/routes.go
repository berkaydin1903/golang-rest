package routes

import (
	"main/controllers"
	"main/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticated)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Get("/api/usercontacts/:id", controllers.GetUserContactById)
	app.Post("/api/adduserContact", controllers.AddUserContact)
	app.Get("/api/usercontact/:id", controllers.GetUserContact)
	app.Put("/api/products/:id", controllers.UpdateUserContact)
	app.Delete("/api/products/:id", controllers.DeleteUserContact)
}
