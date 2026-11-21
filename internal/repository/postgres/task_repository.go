package postgres

import (
	"context"
	"fmt"

	"task-management-go-postgres/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	DB *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		DB: db,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task models.Task) error {
	query := `
		INSERT INTO tasks (title, description, status)
		VALUES ($1, $2, $3)
	`

	_, err := r.DB.Exec(ctx, query,
		task.Title,
		task.Description,
		task.Status,
	)

	return err
}



func (r *TaskRepository) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	query := `
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
		ORDER BY id
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}



func (r *TaskRepository) GetTaskByID(ctx context.Context, id int) (models.Task, error) {
	var task models.Task

	query := `
		SELECT id, title, description, status, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	return task, err
}


func (r *TaskRepository) UpdateTask(ctx context.Context, id int, task models.Task) error {
	query := `
		UPDATE tasks
		SET title = $1,
		    description = $2,
		    status = $3,
		    updated_at = CURRENT_TIMESTAMP
		WHERE id = $4
	`

	result, err := r.DB.Exec(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Status,
		id,
	)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}



func (r *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := r.DB.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("task not found")
	}

	return nil
}
