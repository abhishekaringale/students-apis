# Students API (Go + SQLite)

A clean, production-ready RESTful API built in Go for managing student records. This project is built step-by-step to demonstrate standard Go web development patterns, structured logging, request validation, database connection, and graceful shutdown.

---

## 🚀 Features

- **CGO-Free SQLite Database**: Uses `glebarez/go-sqlite` (a pure Go SQLite driver) so it compiles and runs on any operating system (including Windows) without needing GCC/C compilers.
- **Dependency Injection & Interface Pattern**: The database layer is decoupled from handlers using a `Storage` interface, making it easy to swap databases (e.g., SQLite to PostgreSQL) or use mocks for testing.
- **Request Validation**: Automatically validates incoming request bodies (e.g., checking if email is unique/formatted, required fields) using `go-playground/validator/v10`.
- **Structured Logging**: Uses Go's native `log/slog` for structured, key-value logging.
- **Configuration Management**: Loads environment settings from a YAML config file or environment variables using `cleanenv`.
- **Graceful Shutdown**: Listens for OS interrupt signals (`Ctrl+C`, `SIGINT`, `SIGTERM`) and shuts down the HTTP server cleanly without dropping active connections.

---

## 📁 Project Structure

```text
├── cmd/
│   └── students-api/
│       └── main.go         # Application entry point, router, and server setup
├── config/
│   └── local.yaml          # Local environment configurations (ignored in git)
├── internal/
│   ├── config/
│   │   └── config.go       # Config loader logic
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go # HTTP handler functions (Create, GetById, GetList)
│   ├── types/
│   │   └── types.go        # Student struct models and validation tags
│   └── utils/
│       └── response/
│           └── response.go # Structured JSON response and error helpers
├── storage/
│   ├── storage.go          # Database Storage interface definition
│   └── sqlite/
│       └── sqlite.go       # SQLite implementation of the Storage interface
├── go.mod                  # Package dependencies
└── README.md               # Project documentation
```

---

## 🛠️ Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/health` | Health Check |
| `POST` | `/api/students` | Create a new student (validates request body) |
| `GET` | `/api/students/{id}` | Fetch a student by their ID |
| `GET` | `/api/students` | Fetch a list of all students |

---

## 🏃 How to Run

### Prerequisites
- Go installed (version **1.22+** is required for native path routing).

### Running the App
1. Initialize dependencies:
   ```bash
   go mod tidy
   ```
2. Start the server (pointing to your local config file):
   ```bash
   go run cmd/students-api/main.go -config config/local.yaml
   ```

---

## 🧪 Testing the API (Examples)

### 1. Create a Student (POST)
**PowerShell:**
```powershell
Invoke-RestMethod -Uri "http://localhost:8082/api/students" -Method Post -ContentType "application/json" -Body '{"name": "Abhishek", "email": "abhishek@example.com", "age": 21}'
```
**cURL:**
```bash
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name": "Abhishek", "email": "abhishek@example.com", "age": 21}'
```
*Response:* `{"id":1}`

### 2. Get All Students (GET)
**PowerShell:**
```powershell
Invoke-RestMethod -Uri "http://localhost:8082/api/students" -Method Get
```
*Response:* `[{"id":1,"name":"Abhishek","email":"abhishek@example.com","age":21}]`

### 3. Get Student by ID (GET)
**PowerShell:**
```powershell
Invoke-RestMethod -Uri "http://localhost:8082/api/students/1" -Method Get
```
*Response:* `{"id":1,"name":"Abhishek","email":"abhishek@example.com","age":21}`
