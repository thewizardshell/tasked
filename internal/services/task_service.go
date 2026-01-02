package services

import (
	"context"
	"tasked/internal/domain"
	"tasked/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) GetTaskById(ctx context.Context, id int64) (*domain.Task, error) {
	return s.repo.GetTaskById(ctx, id)
}

func (s *TaskService) ListTaskByUser(ctx context.Context, userId int64) ([]domain.Task, error) {
	return s.repo.ListTaskByUser(ctx, userId)
}

func (s *TaskService) UpdateTask(ctx context.Context, id int64, title string, description string, status string, priority string, dueDate string) (*domain.Task, error) {
	return s.repo.UpdateTask(ctx, id, title, description, status, priority, dueDate)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	return s.repo.DeleteTask(ctx, id)
}

func (s *TaskService) UpdateStatus(ctx context.Context, id int64, status string) (*domain.Task, error) {
	return s.repo.UpdateStatus(ctx, id, status)
}

func (s *TaskService) CreateTask(ctx context.Context, title string, description string, status string, priority string, userId int64, dueDate string) (*domain.Task, error) {
	return s.repo.CreateTask(ctx, title, description, status, priority, userId, dueDate)
}
