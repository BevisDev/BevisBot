package service

import (
	"context"
	"time"

	"github.com/BevisDev/BevisBot/internal/model/entity"
	"github.com/BevisDev/BevisBot/internal/model/request"
	"github.com/BevisDev/BevisBot/internal/model/response"
	"github.com/BevisDev/BevisBot/internal/repository"
)

type TaskService interface {
	Create(ctx context.Context, r *request.Task) (int64, error)
	Search(ctx context.Context, f *request.TaskSearch) ([]*response.Task, int64, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(
	repo repository.TaskRepository,
) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (t *taskService) Create(ctx context.Context, r *request.Task) (int64, error) {
	now := time.Now()
	e := &entity.Task{
		Title:       r.Title,
		Description: r.Description,
		Priority:    r.Priority,
		DueDate:     r.DueDate,
		CreatedBy:   &r.CreatedBy,
		CreatedAt:   &now,
	}
	return t.repo.Create(ctx, e)
}

func (t *taskService) GetTasks(ctx context.Context, r *request.Task) (int64, error) {
	now := time.Now()
	e := &entity.Task{
		Title:       r.Title,
		Description: r.Description,
		Priority:    r.Priority,
		DueDate:     r.DueDate,
		CreatedBy:   &r.CreatedBy,
		CreatedAt:   &now,
	}
	return t.repo.Create(ctx, e)
}

func (t *taskService) Search(ctx context.Context, f *request.TaskSearch) ([]*response.Task, int64, error) {
	return nil, 0, nil
}
