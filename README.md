# TODO-APP

TODO-APP is a simple HTTP API built with Go that integrates with MongoDB to manage todo items. This repository contains the backend implementation of the API.

## Features

- **Create** a new todo item.
- **Read** existing todo items.
- **Update** todo items.
- **Delete** todo items.

## Technologies Used

- **Go** - Programming language used to build the API.
- **MongoDB** - Database used to store todo items.
- **MongoDB Go Driver** - Official MongoDB driver for Go.

## Prerequisites

Before running this project locally, make sure you have the following installed:

- Go
- MongoDB (running locally or accessible via URL)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/lazarpetrovicc/Todo-App.git
   cd Todo-App/API
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Set up MongoDB:

   - Install MongoDB locally or use a cloud-based MongoDB service.
   - Update `mongo.go` file with your MongoDB connection string.

4. Build and run the project:

   ```bash
   go run .
   ```

## API Endpoints

### Get All Todos

- **Method:** GET
- **URL:** `http://localhost:8080/todos`

### Create Todo

- **Method:** POST
- **URL:** `http://localhost:8080/todos`
- **Body:**
  ```json
  {
    "title": "Test Todo",
    "completed": true
  }
  ```

### Get Todo by ID

- **Method:** GET
- **URL:** `http://localhost:8080/todos/{id}`
- **Replace** `{id}` **with a valid todo ID.**

### Update Todo

- **Method:** PUT
- **URL:** `http://localhost:8080/todos/{id}`
- **Body:**
  ```json
  {
    "title": "Updated Todo",
    "completed": true
  }
  ```
- **Replace** `{id}` **with a valid todo ID.**

### Delete Todo

- **Method:** DELETE
- **URL:** `http://localhost:8080/todos/{id}`
- **Replace** `{id}` **with a valid todo ID.**

## Postman Collection

You can find a Postman collection (`Todo-App.postman_collection.json`) in the root of this repository. Import this collection into Postman to easily test the API endpoints.

### Steps to Use Postman Collection

1. Open Postman.
2. Click on `Import` button.
3. Drag and drop `Todo-App.postman_collection.json` file into the window or click `Choose Files` to select it.
4. Once imported, you will see the "Todo-App" collection in the left sidebar of Postman.
5. Expand the collection to see individual requests like "Get All Todos", "Create Todo", etc.
6. Click on any request to open it.
7. Adjust the request URL or body parameters as needed.
8. Click `Send` to make the request to your local API server (make sure your server is running).

## Contributing

Contributions are welcome! If you have any ideas, suggestions, or improvements, please open an issue or submit a pull request.
