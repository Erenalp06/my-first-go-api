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

// CreateTodo godoc
// @Summary Add a new todo
// @Description Add a new todo to the list
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Create Todo"
// @Success 201 {object} models.Todo
// @Router /api/todo [post]
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

// GetAllTodo godoc
// @Summary Get all todos
// @Description Get a list of all todos
// @Tags todos
// @Produce json
// @Success 200 {array} models.Todo
// @Router /api/todo [get]
func (h TodoHandler) GetAllTodo(c *fiber.Ctx) error {
	result, err := h.Service.GetAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(result)
}

// DeleteTodoByID godoc
// @Summary Delete a todo by ID
// @Description Delete a single todo by its ID
// @Tags todos
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} map[string]bool "State: true if deleted"
// @Router /api/todo/{id} [delete]
func (h TodoHandler) DeleteTodoByID(c *fiber.Ctx) error {
	query := c.Params("id")
	cnv, _ := primitive.ObjectIDFromHex(query)

	result, err := h.Service.Delete(cnv)

	if err != nil || !result {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"State": false})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"State": true})
}

// GetTodoById godoc
// @Summary Get a todo by ID
// @Description Get details of a single todo by its ID
// @Tags todos
// @Produce json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.Todo
// @Router /api/todo/{id} [get]
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

// GetTodoByTitle godoc
// @Summary Get todos by title
// @Description Get todos that match a specific title
// @Tags todos
// @Produce json
// @Param title path string true "Todo Title"
// @Success 200 {array} models.Todo
// @Router /api/todo/title/{title} [get]
func (h TodoHandler) GetTodoByTitle(c *fiber.Ctx) error {
	query := c.Params("title")

	result, err := h.Service.GetByTitle(query)
	if err != nil {
		log.Printf("%v", err.Error())
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusOK).JSON(result)
}

// UpdateById godoc
// @Summary Update a todo by ID
// @Description Update details of a todo by its ID
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "Todo ID"
// @Param todo body models.Todo true "Update Todo"
// @Success 200 {object} map[string]string "state: record was updated successfully"
// @Router /api/todo/{id} [put]
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
