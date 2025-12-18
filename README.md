Go - Backend Development Task
🧪 User API with DOB, Age Calculation & Clean Architecture

This project is a clean, modular RESTful API built using Go, Fiber, SQLC, PostgreSQL, and Zap logging.
It manages users with name and dob, and dynamically calculates age using Go’s time package.

📂 Project Structure
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

🔧 Tech Stack

GoFiber

PostgreSQL

SQLC

Uber Zap

go-playground/validator

JWT Authentication

🗄️ Database Schema
users table
Field	Type	Constraints
id	SERIAL	PRIMARY KEY
name	TEXT	NOT NULL
dob	DATE	NOT NULL
auth_users table
Field	Type	Constraints
id	SERIAL	PRIMARY KEY
name	TEXT	NOT NULL
email	TEXT	UNIQUE NOT NULL
password_hash	TEXT	NOT NULL
🔄 API Endpoints
➕ Create User

POST /users

Request

{
  "name": "Alice",
  "dob": "1990-05-10"
}


Response

{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10"
}

🔍 Get User By ID

GET /users/:id

Response

{
  "id": 1,
  "name": "Alice",
  "dob": "1990-05-10",
  "age": 35
}

✏️ Update User

PUT /users/:id

Request

{
  "name": "Alice Updated",
  "dob": "1991-03-15"
}


Response

{
  "id": 1,
  "name": "Alice Updated",
  "dob": "1991-03-15"
}

🗑 Delete User

DELETE /users/:id

Response:

204 No Content

📃 List All Users

GET /users

Response

[
  {
    "id": 1,
    "name": "Alice",
    "dob": "1990-05-10",
    "age": 34
  }
]

🔐 Authentication Endpoints
Register

POST /auth/register

{
  "name": "Ravi",
  "email": "ravi@gmail.com",
  "password": "password123"
}

Login

POST /auth/login

{
  "email": "ravi@gmail.com",
  "password": "password123"
}


Response

{
  "token": "jwt_token_here"
}

Get Current User

GET /auth/me

Headers

Authorization: Bearer <token>


Response

{
  "id": 1,
  "name": "Ravi",
  "email": "ravi@gmail.com"
}

🧮 Age Calculation Logic
func CalculateAge(dob time.Time) int {
    now := time.Now()
    age := now.Year() - dob.Year()

    if now.Month() < dob.Month() || 
       (now.Month() == dob.Month() && now.Day() < dob.Day()) {
        age--
    }
    return age
}

🚀 Run the Project
1️⃣ Clone the repo
git clone <repo-url>
cd <project-folder>

2️⃣ Install dependencies
go mod tidy

3️⃣ Setup environment

Create .env:

DATABASE_URL=postgres://user:password@localhost:5432/db?sslmode=disable
JWT_SECRET=supersecretkey

4️⃣ Run migrations
migrate -path db/migrations -database "$DATABASE_URL" up

5️⃣ Generate SQLC code
sqlc generate

6️⃣ Start server
go run cmd/server/main.go


Server runs at:
👉 http://localhost:3000

🧪 Validation & Logging

go-playground/validator used for input validation

Zap logger used for structured logs

Clean error messages + proper HTTP status codes

📦 Optional Enhancements

Docker support

Pagination

Unit test for age calculation

Request ID middleware

Request duration logging

📝 Submission Checklist

✔ Push to GitHub
✔ Include README.md
✔ Add .env example (no secrets)
✔ Share the repo link

🎉 Project Completed

Your backend meets all assignment requirements and includes extra professional features.