package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func SwaggerRoute(app *fiber.App) {

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	err := app.Listen(":8087")
	if err != nil {
		log.Fatalf("fiver.Listen failed %s", err)
	}

}
