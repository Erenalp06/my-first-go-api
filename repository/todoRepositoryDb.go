package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Erenalp06/first-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
	GetAll() ([]models.Todo, error)
	Delete(id primitive.ObjectID) (bool, error)
	GetById(id primitive.ObjectID) (models.Todo, error)
	GetByTitle(title string) (models.Todo, error)
	Update(id primitive.ObjectID, updatedTodo models.Todo) (bool, error)
}

func (t *TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	todo.Id = primitive.NewObjectID()

	result, err := t.TodoCollection.InsertOne(ctx, todo)

	if result.InsertedID == nil || err != nil {
		errors.New("failed add")
		return false, err
	}

	return true, nil
}

func (t *TodoRepositoryDB) GetAll() ([]models.Todo, error) {
	var todo models.Todo
	var todos []models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.Find(ctx, bson.M{})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	for result.Next(ctx) {
		if err := result.Decode(&todo); err != nil {
			log.Fatalln(err)
			return nil, err
		}

		todos = append(todos, todo)
	}
	return todos, nil
}

func (t *TodoRepositoryDB) GetById(id primitive.ObjectID) (models.Todo, error) {
	var todo models.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := t.TodoCollection.FindOne(ctx, bson.M{"id": id}).Decode(&todo)

	if err != nil {

		if err == mongo.ErrNoDocuments {
			log.Printf("No document")
			return models.Todo{}, err
		}

		log.Printf("Error finding todo in the database: %v, Error: %v\n", id, err)
		return models.Todo{}, err
	}

	return todo, nil

}

func (t *TodoRepositoryDB) Delete(id primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := t.TodoCollection.DeleteOne(ctx, bson.M{"id": id})

	if err != nil || result.DeletedCount <= 0 {
		return false, err
	}

	return true, nil
}

func (t *TodoRepositoryDB) GetByTitle(title string) (models.Todo, error){
	var todo models.Todo
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := t.TodoCollection.FindOne(ctx, bson.M{"title": title}).Decode(&todo)

	if err!=nil{
		if err == mongo.ErrNoDocuments{
			log.Printf("No document with given title : %v", title)
			return models.Todo{}, err
		}

		log.Printf("Error finding todo in the database: %v, Error: %v\n", title, err)
		return models.Todo{}, err		
	}

	return todo, nil

}

func (t *TodoRepositoryDB) Update(id primitive.ObjectID, updatedTodo models.Todo) (bool, error){
	ctx,cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updatedTodo.Id = id

	updateResult, err := t.TodoCollection.UpdateOne(
		ctx,
		bson.M{"id": id},
		bson.M{"$set": updatedTodo},
	)

	if err!=nil{
		log.Printf("Error updating todo in the database: %v, Error: %v\n",id,  err.Error())
		return false,err
	}

	if updateResult.MatchedCount == 0{
		return false, errors.New("no todo found with given ID")
	}

	return true, nil

}

func NewToDoRepositoryDb(dbClient *mongo.Collection) *TodoRepositoryDB {
	return &TodoRepositoryDB{TodoCollection: dbClient}
}
