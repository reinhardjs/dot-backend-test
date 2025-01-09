# Dot Backend Test

## Overview

This project is a backend application built using Go, GORM, and Gin. It provides a RESTful API for managing products and categories, with support for CRUD operations and caching using Redis.

## Features

- **Product Management**: Create, read, update, and delete products.
- **Category Management**: Create, read, update, and delete categories.
- **Caching**: Utilizes Redis for caching category and product data to improve performance.
- **Database**: Uses PostgreSQL as the primary database.

## Technologies Used

- Go (1.22.0)
- Gin (v1.10.0)
- GORM (v1.25.12)
- PostgreSQL (v10.4)
- Redis
- Docker

## Getting Started

### Prerequisites

- Go installed on your machine.
- Docker and Docker Compose for running the database and Redis.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/reinhardjs/dot-backend-test.git
   cd dot-backend-test
   ```

2. Create a `.env` file in the root directory and configure your environment variables as needed.

3. Start the services using Docker Compose:

   ```bash
   docker-compose up -d
   ```

4. Run the application:

   ```bash
   go run cmd/main.go
   ```

### Building Go Application

To build the Go application, run the following command in the project directory:

```bash
go build -o main cmd/main.go
```

### Building Docker Image

To build the Docker image, run the following command in the project directory:

```bash
docker build -t dot-backend-test .
```

### Running Docker Container

To run the Docker container, use the following command:

```bash
docker run -p 8080:8080 dot-backend-test
```

### API Endpoints and Example Payloads

#### Products

- **Create Product**
  - `POST /api/v1/products`
  - Request Body Example:
    ```json
    {
      "name": "Sample Product",
      "price": 29.99,
      "category_id": 1
    }
    ```
  - Response (201 Created):
    ```json
    {
      "id": 1,
      "name": "Sample Product",
      "price": 29.99,
      "category_id": 1,
      "created_at": "2024-03-14T12:00:00Z",
      "updated_at": "2024-03-14T12:00:00Z"
    }
    ```

- **Get All Products**
  - `GET /api/v1/products`
  - Response (200 OK):
    ```json
    [
      {
        "id": 1,
        "name": "Sample Product",
        "price": 29.99,
        "category_id": 1,
        "created_at": "2024-03-14T12:00:00Z",
        "updated_at": "2024-03-14T12:00:00Z"
      }
    ]
    ```
  
- **Get Product by ID**
  - `GET /api/v1/products/:id`
  - Response (200 OK):
    ```json
    {
      "id": 1,
      "name": "Sample Product",
      "price": 29.99,
      "category_id": 1,
      "created_at": "2024-03-14T12:00:00Z",
      "updated_at": "2024-03-14T12:00:00Z"
    }
    ```

- **Update Product**
  - `PUT /api/v1/products/:id`
  - Request Body Example:
    ```json
    {
      "name": "Updated Product",
      "price": 39.99,
      "category_id": 1
    }
    ```
  - Response (200 OK):
    ```json
    {
      "id": 1,
      "name": "Updated Product",
      "price": 39.99,
      "category_id": 1,
      "created_at": "2024-03-14T12:00:00Z",
      "updated_at": "2024-03-14T12:30:00Z"
    }
    ```

- **Delete Product**
  - `DELETE /api/v1/products/:id`
  - Response (200 OK):
    ```json
    {
      "message": "Product deleted successfully"
    }
    ```

#### Categories

- **Create Category**
  - `POST /api/v1/categories`
  - Request Body Example:
    ```json
    {
      "name": "Electronics"
    }
    ```
  - Response (201 Created):
    ```json
    {
      "id": 1,
      "name": "Electronics",
      "created_at": "2024-03-14T12:00:00Z",
      "updated_at": "2024-03-14T12:00:00Z"
    }
    ```

- **Get All Categories**
  - `GET /api/v1/categories`
  - Response (200 OK):
    ```json
    [
      {
        "id": 1,
        "name": "Electronics",
        "created_at": "2024-03-14T12:00:00Z",
        "updated_at": "2024-03-14T12:00:00Z"
      }
    ]
    ```

- **Get Category by ID**
  - `GET /api/v1/categories/:id`
  - Response (200 OK):
    ```json
    {
      "id": 1,
      "name": "Electronics",
      "created_at": "2024-03-14T12:00:00Z",
      "updated_at": "2024-03-14T12:00:00Z"
    }
    ```

- **Update Category**
  - `PUT /api/v1/categories/:id`
  - Request Body Example:
    ```json
    {
      "name": "Updated Electronics"
    }
    ```
  - Response (200 OK):
    ```json
    {
      "id": 1,
      "name": "Updated Electronics",
      "created_at": "2024-03-14T12:00:00Z",
      "updated_at": "2024-03-14T12:30:00Z"
    }
    ```

- **Delete Category**
  - `DELETE /api/v1/categories/:id`
  - Response (200 OK):
    ```json
    {
      "message": "Category deleted successfully"
    }
    ```

## Running Tests

### Go to test directory
To run the tests, use the following command:

```
cd test
```

### Start the test environment:
```
docker compose up -d
```

### Run the test

```
go test ./... -v
```

The e2e tests will automatically:
   - Clean the database
   - Create test categories
   - Create test products
   - Verify CRUD operations

