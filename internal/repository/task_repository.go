package repository

import (
	"context"

	"github.com/BevisDev/BevisBot/internal/lib"
	"github.com/BevisDev/BevisBot/internal/model/entity"
	"github.com/BevisDev/BevisBot/internal/model/request"
)

type TaskRepository interface {
	Create(ctx context.Context, task *entity.Task) (int64, error)
	Search(ctx context.Context, f *request.TaskSearch) ([]*entity.Task, int64, error)
}

type taskRepository struct {
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{}
}

func (t *taskRepository) Create(ctx context.Context, task *entity.Task) (int64, error) {
	if err := lib.DB.WithContext(ctx).
		Create(task).
		Error; err != nil {
		return 0, err
	}
	return task.Id, nil
}

func (t *taskRepository) Search(ctx context.Context,
	f *request.TaskSearch,
) ([]*entity.Task, int64, error) {
	return nil, 0, nil
}

func (t *taskRepository) Tasks(ctx context.Context) (*entity.Task, error) {
	return nil, nil
}
