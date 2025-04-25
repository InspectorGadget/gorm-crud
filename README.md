# gorm-crud (Gin Gonic + Gorm + MySQL)

## Description
This is a simple CRUD application built with Go, using the Gin Gonic framework for routing and Gorm for ORM. The application connects to a MySQL database and provides basic CRUD operations.

## Features
- Create, Read, Update, Delete operations
- RESTful API
- MySQL database integration
- Gorm ORM for database interactions

## Prerequisites
- Go 1.16 or later
- Docker or MySQL installed

## Guide
### 1. Clone the repository
```bash
git clone https://github.com/InspectorGadget/gorm-crud.git
cd gorm-crud
```

### 2. Set up the database
```bash
docker compose up -d
```
```bash
go run migrate/migration.go
```

### 3. Setup the Go environment
```bash
go get
go mod tidy
```

### 4. Run the application
```bash
go run .
```

### 5. Test the API
You can use Postman or curl to test the API endpoints or use the built-in HTTP API Testkit in `./api-requests` folder.