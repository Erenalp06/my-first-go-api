# Todo API with Go Fiber and Swagger

This is a simple RESTful API for managing todo items, implemented in Go using the Fiber web framework. It provides endpoints for creating, retrieving, updating, and deleting todo items. Swagger is used for API documentation.

## Features

- Create, Read, Update, Delete (CRUD) operations for todo items.
- Swagger UI integration for API documentation and testing.
- MongoDB for persistent storage.

## Installation

Make sure you have Go installed on your system. [Download Go](https://golang.org/dl/)

Clone the repository to your local machine:

```bash
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
```

Install the dependencies:
```bash
go mod tidy
```

## Running the Application
To start the server, run:

```bash
go run main.go
```

## Usage
The following endpoints are available:

- POST /api/todo - Create a new todo item.
- GET /api/todo - Retrieve all todo items.
- GET /api/todo/{id} - Retrieve a todo item by its ID.
- PUT /api/todo/{id} - Update a todo item by its ID.
- DELETE /api/todo/{id} - Delete a todo item by its ID.
- GET /api/todo/title/{title} - Retrieve todo items by title.

## Swagger UI
To access the Swagger UI and interact with the API, visit:
```
http://localhost:8087/swagger/index.html
```

### Create a Todo Item
This section describes how to create a new todo item using the POST method. Each todo item requires a title and content.

#### HTTP Request

* **Method:** POST
* **Endpoint:** /api/todo
* **Headers:**
  * **Content-Type:** application/json

**Request Body**
Provide the title and content for the todo item in the request body as JSON.
```json
"title": "Sample Todo",
"content": "This is a sample todo item."
```

**Successful Response**
A successful request returns the HTTP status code 201 Created and the created todo item, including its generated ID.

* **Status Code:** 201 Created
* **Response Body:**

```json
"id": "5f6f4f4b50956f91ca8892",
"title": "Sample Todo",
"content": "This is a sample todo item."
```





