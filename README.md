# Task Management API using Go and PostgreSQL

A RESTful Task Management API built using **Go**, **PostgreSQL**, and Go's standard `net/http` package. This project supports creating, retrieving, updating, and deleting tasks.

## Features

- Create a new task
- Retrieve all tasks
- Retrieve a task by ID
- Update an existing task
- Delete a task
- PostgreSQL database integration
- Input validation
- Database connection using pgx connection pool
- REST API testing using Postman

## Technologies Used

- **Language:** Go
- **Database:** PostgreSQL
- **HTTP:** Go standard library (`net/http`)
- **Database Driver:** `pgx/v5/pgxpool`
- **Environment Variables:** godotenv
- **API Testing:** Postman
- **IDE:** Visual Studio Code

## Project Structure

```text
task-management-go-postgres/
│
├── internal/
│   ├── handler/
│   │   └── task_handler.go
│   │
│   ├── models/
│   │   └── task.go
│   │
│   ├── repository/
│   │   └── postgres/
│   │       ├── connection.go
│   │       └── task_repository.go
│   │
│   └── service/
│       └── task_service.go
│
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── main.go
```

## Database Setup

1. Install PostgreSQL.
2. Create a database named `task_management`.
3. Execute the following SQL query:

```sql
CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

## Environment Configuration

Create a `.env` file in the project root directory.

Add your PostgreSQL connection string:

```env
DATABASE_URL=postgres://username:password@localhost:5432/task_management
```

Replace `username` and `password` with your PostgreSQL credentials.

**Important:** Never upload your `.env` file or database credentials to GitHub.

## Installation and Execution

Clone the repository:

```bash
git clone https://github.com/suryanarayana1234/task-management-go-postgres.git
```

Navigate to the project directory:

```bash
cd task-management-go-postgres
```

Install dependencies:

```bash
go mod download
```

Run the application:

```bash
go run .
```

Server runs at:

```text
http://localhost:8080
```

## API Endpoints

| Method | Endpoint      | Description         |
| ------ | ------------- | ------------------- |
| POST   | `/tasks`      | Create a task       |
| GET    | `/tasks`      | Get all tasks       |
| GET    | `/tasks/{id}` | Get task by ID      |
| PUT    | `/tasks/{id}` | Update a task       |
| DELETE | `/tasks/{id}` | Delete a task       |
| GET    | `/health`     | Check server status |

## Sample Request

**Create Task**

Method: `POST`

URL: `http://localhost:8080/tasks`

Body (JSON):

```json
{
  "title": "Learn Go",
  "description": "Learn Go with PostgreSQL",
  "status": "pending"
}
```

**Update Task**

Method: `PUT`

URL: `http://localhost:8080/tasks/1`

Body (JSON):

```json
{
  "title": "Learn Advanced Go",
  "description": "Build projects using Go and PostgreSQL",
  "status": "completed"
}
```

## Task Status Values

- `pending`
- `in_progress`
- `completed`

## Testing

All API endpoints can be tested using Postman.

## Author

**Surya Narayana**

GitHub: https://github.com/suryanarayana1234

## License

This project is open-source and available for learning and educational purposes.
