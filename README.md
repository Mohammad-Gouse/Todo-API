# Golang TODO API

This is a Golang-based RESTful API for managing TODO items. The API supports basic CRUD operations and includes pagination and filtering functionality for the list endpoint.

## Features

- Create, Read, Update, and Delete TODO items
- Pagination and filtering for listing TODO items
- Filter TODO items based on status and creation date

## Technologies Used

- Golang
- MongoDB
- Gin Web Framework

## Prerequisites

- Go
- MongoDB
- Gin


## Project Setup

1. **Clone the repository:**
   git clone https://github.com/Mohammad-Gouse/Todo-API.git
   cd todo-api

2. **Install dependencies:**
   go get go.mongodb.org/mongo-driver/mongo
   go get github.com/gin-gonic/gin
   go get github.com/joho/godotenv

3. **Configure MongoDB:**
    Update the MongoDB URI in utils/utils.go:
    const mongoURI = "mongodb://localhost:27017"

4. **Run the application:**
    go run main.go


API Endpoints
Create a TODO
Endpoint: /todos
Method: POST
request body:
{
    "user_id": "user123",
    "title": "Buy groceries",
    "description": "Milk, Bread, Butter",
    "status": "pending"
}


Get TODOs (with Pagination and Filters)
Endpoint: /todos
Method: GET
Query Parameters:
user_id: (required) The ID of the user
status: (optional) The status of TODO items (pending, completed)
page: (optional) The page number (default: 1)
limit: (optional) The number of items per page (default: 10)
created_from: (optional) The start date for the creation date filter
created_to: (optional) The end date for the creation date filter

example:
http://localhost:8080/todos?user_id=user3&status=pending&created_from=2024-07-15T17:52:55.834Z

response:
[
    {
        "id": "669561f709ca84f049dcbca3",
        "user_id": "user3",
        "title": "title5",
        "description": "description",
        "status": "pending",
        "created": "2024-07-15T17:52:55.834Z",
        "updated": "2024-07-15T17:52:55.834Z"
    }
]

Update a TODO
Endpoint: /todos/:id
Method: PUT

example:
url and request body:
http://localhost:8080/todos/6693e7ec1f7821dfcfc27664

{
    "title": "Buy groceries and snack",
    "description": "Milk, Bread, Butter, Chips",
    "status": "pending"
}

response:
{
    "message": "Todo updated successfully"
}

Delete a TODO
Endpoint: /todos/:id
Method: DELETE

example:
http://localhost:8080/todos/6693e7ec1f7821dfcfc27664

respone:
{
    "message": "Todo deleted successfully"
}

