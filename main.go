package main

import (
	"github.com/gofiber/fiber/v2"
	"tdd-go-fiber/controllers"
	"tdd-go-fiber/usecase"
)

func main() {
	app := fiber.New()

	SetupRoutes(app)
	app.Listen(":3000")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/contacts", func(c *fiber.Ctx) error {

		cu := usecase.NewContactUseCase()
		return controllers.GetContacts(c, cu)
	})
}
