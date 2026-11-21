package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"

	"task-management-go-postgres/internal/handler"
	"task-management-go-postgres/internal/repository/postgres"
	"task-management-go-postgres/internal/service"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Connect to PostgreSQL
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db, err := postgres.Connect(ctx)
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	log.Println("PostgreSQL connected successfully!")

	// Initialize layers
	taskRepo := postgres.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// Create HTTP routes
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Task Management API is running"))
	})

	// Create and get all tasks
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			taskHandler.CreateTask(w, r)
		case http.MethodGet:
			taskHandler.GetAllTasks(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Get and update task by ID
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetTaskByID(w, r)
		case http.MethodPut:
			taskHandler.UpdateTask(w, r)
		case http.MethodDelete:
			taskHandler.DeleteTask(w, r)	
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
