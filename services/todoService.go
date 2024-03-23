package services

import (
	"errors"
	"log"

	"github.com/Erenalp06/first-go-api/dto"
	"github.com/Erenalp06/first-go-api/models"
	"github.com/Erenalp06/first-go-api/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DefaultToDoService struct {
	Repo repository.TodoRepository
}

type TodoService interface {
	Insert(todo models.Todo) (*dto.TodoDTO, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
	GetById(id primitive.ObjectID) (models.Todo, error)
	GetByTitle(title string) (models.Todo, error)
	Update(id primitive.ObjectID, updatedTodo models.Todo) (bool, error)
}

func (t *DefaultToDoService) Insert(todo models.Todo) (*dto.TodoDTO, error) {
	var response dto.TodoDTO	
	if len(todo.Title) <= 2 {
		response.Status = false
		return &response, nil
	}

	result, err := t.Repo.Insert(todo)

	if err != nil || result == false {
		response.Status = false
		return &response, err
	}

	response = dto.TodoDTO{Status: result}

	return &response, nil
}

func (t *DefaultToDoService) GetAll() ([]models.Todo, error) {
	result, err := t.Repo.GetAll()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *DefaultToDoService) GetById(id primitive.ObjectID) (models.Todo, error) {
	result, err := t.Repo.GetById(id)
	if err != nil {
		log.Printf("Error fetching todo by ID: %v, Error: %v\n", id, err)
		return models.Todo{}, err

	}
	return result, nil
}

func (t *DefaultToDoService) GetByTitle(title string) (models.Todo, error) {
	result, err := t.Repo.GetByTitle(title)
	if err != nil {
		log.Print(err.Error())
		return models.Todo{}, err
	}
	return result, nil
}

func (t *DefaultToDoService) Delete(id primitive.ObjectID) (bool, error) {
	result, err := t.Repo.Delete(id)
	if err != nil || !result {
		return false, err
	}
	return true, nil
}

func (t *DefaultToDoService) Update(id primitive.ObjectID, updatedTodo models.Todo) (bool, error) {

	if updatedTodo.Title == "" || updatedTodo.Content == "" {
		errMsg := "title and content must not be empty"
		log.Print(errMsg)
		return false, errors.New(errMsg)
	}

	result, err := t.Repo.Update(id, updatedTodo)
	if err != nil || !result {
		log.Print(err.Error())
		return false, err
	}
	return true, nil
}

func NewTodoService(repo repository.TodoRepository) *DefaultToDoService {
	return &DefaultToDoService{Repo: repo}
}
