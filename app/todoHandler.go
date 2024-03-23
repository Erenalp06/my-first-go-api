package app

import (
	"log"
	"net/http"

	"github.com/Erenalp06/first-go-api/models"
	"github.com/Erenalp06/first-go-api/services"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandler struct {
	Service services.TodoService
}

func (h TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	result, err := h.Service.Insert(todo)

	if err != nil || result.Status == false {
		return err
	}

	return c.Status(http.StatusCreated).JSON(result)
}

func (h TodoHandler) GetAllTodo(c *fiber.Ctx) error {
	result, err := h.Service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h TodoHandler) DeleteTodoByID(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.Delete(cnv)

	if err != nil || !result {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}

func (h TodoHandler) GetTodoById(c *fiber.Ctx) error {
	query := c.Params("id")
	convert, err := primitive.ObjectIDFromHex(query)
	if err != nil {
		// ID dönüştürme hatası
		log.Printf("%v", err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}

	result, err := h.Service.GetById(convert)
	if err != nil {
		log.Printf("Error fetching todo by ID: %v, Error: %v\n", convert, err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if result.Id.IsZero() {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "no todo found with the given ID"})
	}

	return c.Status(http.StatusOK).JSON(result)
}

func (h TodoHandler) GetTodoByTitle(c *fiber.Ctx) error {
	query := c.Params("title")

	result, err := h.Service.GetByTitle(query)
	if err != nil {
		log.Printf("%v", err.Error())
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}

func (h TodoHandler) UpdateById(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	}

	query := c.Params("id")
	convert, err := primitive.ObjectIDFromHex(query)
	if err != nil {
		log.Print(err.Error())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	result, err := h.Service.Update(convert, todo)
	if err != nil || !result {
		log.Print(err.Error())
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"state": "record was updated successfuly"})
}
