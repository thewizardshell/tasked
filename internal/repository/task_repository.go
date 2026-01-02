package repository

import (
	"context"
	"database/sql"
	"fmt"
	"tasked/internal/database"
	"tasked/internal/domain"
	"time"
)

type TaskRepository interface {
	GetTaskById(ctx context.Context, id int64) (*domain.Task, error)
	ListTaskByUser(ctx context.Context, id int64) ([]domain.Task, error)
	UpdateTask(ctx context.Context, id int64, title string, description string, status string, priority string, dueDate string) (*domain.Task, error)
	DeleteTask(ctx context.Context, id int64) error
	UpdateStatus(ctx context.Context, id int64, status string) (*domain.Task, error)
	CreateTask(ctx context.Context, title string, description string, status string, priority string, userId int64, dueDate string) (*domain.Task, error)
}

type taskRepository struct {
	queries *database.Queries
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		queries: database.New(db),
	}
}

func (r *taskRepository) GetTaskById(ctx context.Context, id int64) (*domain.Task, error) {
	dbTask, err := r.queries.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &domain.Task{
		Id:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description.String,
		Status:      dbTask.Status.String,
		Priority:    dbTask.Priority.String,
		Userid:      dbTask.UserID,
		Duedate:     dbTask.DueDate.Time,
		CompletedAt: dbTask.CompletedAt.Time,
		UpdatedAt:   dbTask.UpdatedAt.Time,
	}, nil
}

func (r *taskRepository) ListTaskByUser(ctx context.Context, id int64) ([]domain.Task, error) {
	dbTasks, err := r.queries.ListTasksByUser(ctx, id)
	if err != nil {
		return nil, err
	}
	tasks := make([]domain.Task, 0, len(dbTasks))
	for _, t := range dbTasks {
		task := domain.Task{
			Id:          t.ID,
			Title:       t.Title,
			Description: t.Description.String,
			Status:      t.Status.String,
			Priority:    t.Priority.String,
			Userid:      t.UserID,
			Duedate:     t.DueDate.Time,
			CompletedAt: t.CompletedAt.Time,
			UpdatedAt:   t.UpdatedAt.Time,
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepository) UpdateTask(ctx context.Context, id int64, title string, description string, status string, priority string, dueDate string) (*domain.Task, error) {
	var nullDueDate sql.NullTime
	if dueDate != "" {
		parsedDate, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return nil, fmt.Errorf("formato de fecha inválido: %w", err)
		}
		nullDueDate = sql.NullTime{
			Time:  parsedDate,
			Valid: true,
		}
	}

	dbTask, err := r.queries.UpdateTask(ctx, database.UpdateTaskParams{
		ID:    id,
		Title: title,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		Status: sql.NullString{
			String: status,
			Valid:  status != "",
		},
		Priority: sql.NullString{
			String: priority,
			Valid:  priority != "",
		},
		DueDate: nullDueDate,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		Id:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description.String,
		Status:      dbTask.Status.String,
		Priority:    dbTask.Priority.String,
		Userid:      dbTask.UserID,
		Duedate:     dbTask.DueDate.Time,
		CompletedAt: dbTask.CompletedAt.Time,
		UpdatedAt:   dbTask.UpdatedAt.Time,
	}, nil
}

func (r *taskRepository) DeleteTask(ctx context.Context, id int64) error {
	return r.queries.DeleteTask(ctx, id)
}

func (r *taskRepository) UpdateStatus(ctx context.Context, id int64, status string) (*domain.Task, error) {
	dbTask, err := r.queries.UpdateTaskStatus(ctx, database.UpdateTaskStatusParams{
		ID: id,
		Status: sql.NullString{
			String: status,
			Valid:  status != "",
		},
	})
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		Id:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description.String,
		Status:      dbTask.Status.String,
		Priority:    dbTask.Priority.String,
		Userid:      dbTask.UserID,
		Duedate:     dbTask.DueDate.Time,
		CompletedAt: dbTask.CompletedAt.Time,
		UpdatedAt:   dbTask.UpdatedAt.Time,
	}, nil
}

func (r *taskRepository) CreateTask(ctx context.Context, title string, description string, status string, priority string, userId int64, dueDate string) (*domain.Task, error) {
	var nullDueDate sql.NullTime
	if dueDate != "" {
		parsedDate, err := time.Parse("2006-01-02", dueDate)
		if err != nil {
			return nil, fmt.Errorf("formato de fecha inválido: %w", err)
		}
		nullDueDate = sql.NullTime{
			Time:  parsedDate,
			Valid: true,
		}
	}

	dbTask, err := r.queries.CreateTask(ctx, database.CreateTaskParams{
		Title: title,
		Description: sql.NullString{
			String: description,
			Valid:  description != "",
		},
		Status: sql.NullString{
			String: status,
			Valid:  status != "",
		},
		Priority: sql.NullString{
			String: priority,
			Valid:  priority != "",
		},
		UserID:  userId,
		DueDate: nullDueDate,
	})
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		Id:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description.String,
		Status:      dbTask.Status.String,
		Priority:    dbTask.Priority.String,
		Userid:      dbTask.UserID,
		Duedate:     dbTask.DueDate.Time,
		CompletedAt: dbTask.CompletedAt.Time,
		UpdatedAt:   dbTask.UpdatedAt.Time,
	}, nil
}
