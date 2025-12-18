## Go - Backend Development Task
User API with DOB, Age Calculation & Clean Architecture

This project is a clean and modular RESTful API built using Go, Fiber, SQLC, PostgreSQL, and Zap logging.
It manages users with name and dob, and dynamically calculates age using Go’s time package.

# Project Structure
```
/cmd/server/main.go
/config/
/db/
 ├── migrations/
 └── sqlc/
/internal/
 ├── handler/
 ├── repository/
 ├── service/
 ├── routes/
 ├── middleware/
 ├── models/
 └── logger/
.env
sqlc.yaml
go.mod

```
# Tech Stack
```
GoFiber
PostgreSQL
SQLC
Uber Zap
go-playground/validator
JWT Authentication

```
# Run the Project

``` git clone https://github.com/Ravi-Pr4kash/ainyx ```
``` cd ainyx ```

## Install dependencies
``` go mod tidy ```

## Create .env
``` DATABASE_URL=postgres://user:password@localhost:5432/db?sslmode=disable ```
```JWT_SECRET=supersecretkey ```

## Run migrations
``` migrate -path db/migrations -database "$DATABASE_URL" up ```
## Generate SQLC code
``` sqlc generate ```
## Start the server
``` go run cmd/server/main.go ```


