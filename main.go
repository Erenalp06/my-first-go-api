package main

import (
	"github.com/Erenalp06/first-go-api/configs"
	"github.com/Erenalp06/first-go-api/repository"
	"github.com/Erenalp06/first-go-api/routes"
	"github.com/Erenalp06/first-go-api/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	configs.SetupLogger()
	appRoute := fiber.New()
	configs.ConnectDB()
	dbClient := configs.GetCollection(configs.DB, "todos")

	TodoRepositoryDb := repository.NewToDoRepositoryDb(dbClient)
	todoService := services.NewTodoService(TodoRepositoryDb)

	routes.SetupRoutes(appRoute, todoService)

	appRoute.Listen(":8087")
}
