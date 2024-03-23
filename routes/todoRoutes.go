package routes

import (
	"github.com/Erenalp06/first-go-api/app"
	_ "github.com/Erenalp06/first-go-api/docs"
	"github.com/Erenalp06/first-go-api/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(appRoute *fiber.App, todoService services.TodoService) {
	td := app.TodoHandler{Service: todoService}

	appRoute.Post("/api/todo", td.CreateTodo)
	appRoute.Get("/api/todo", td.GetAllTodo)
	appRoute.Delete("/api/todo/:id", td.DeleteTodoByID)
	appRoute.Get("/api/todo/:id", td.GetTodoById)
	appRoute.Get("/api/todo/title/:title", td.GetTodoByTitle)
	appRoute.Put("/api/todo/:id", td.UpdateById)
}
