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

### API Endpoints

#### Products

- **Create Product**
  - `POST /api/v1/products`
  - Request Body: JSON representation of the product.
  
- **Get All Products**
  - `GET /api/v1/products`
  
- **Get Product by ID**
  - `GET /api/v1/products/:id`
  
- **Update Product**
  - `PUT /api/v1/products/:id`
  - Request Body: JSON representation of the product.
  
- **Delete Product**
  - `DELETE /api/v1/products/:id`

#### Categories

- **Create Category**
  - `POST /api/v1/categories`
  - Request Body: JSON representation of the category.
  
- **Get All Categories**
  - `GET /api/v1/categories`
  
- **Get Category by ID**
  - `GET /api/v1/categories/:id`
  
- **Update Category**
  - `PUT /api/v1/categories/:id`
  - Request Body: JSON representation of the category.
  
- **Delete Category**
  - `DELETE /api/v1/categories/:id`

## Running Tests

To run the tests, use the following command:
