package service

import (
	"context"

	"task-management-go-postgres/internal/models"
	"task-management-go-postgres/internal/repository/postgres"
)

type TaskService struct {
	repo *postgres.TaskRepository
}

func NewTaskService(repo *postgres.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, task models.Task) error {
	return s.repo.CreateTask(ctx, task)
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetAllTasks(ctx)
}


func (s *TaskService) GetTaskByID(ctx context.Context, id int) (models.Task, error) {
	return s.repo.GetTaskByID(ctx, id)
}


func (s *TaskService) UpdateTask(ctx context.Context, id int, task models.Task) error {
	return s.repo.UpdateTask(ctx, id, task)
}


func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}
