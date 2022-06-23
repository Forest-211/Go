package routes

import (
	"ambassador/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("api/v1")

	admin := api.Group("admin")
	admin.Post("/register", controllers.Register)
}