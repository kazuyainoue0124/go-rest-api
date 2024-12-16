package domain

import "context"

type ITaskRepository interface {
	GetAllTasks(ctx context.Context) ([]*Task, error)
	GetTaskById(ctx context.Context, id int64) (*Task, error)
	CreateTask(ctx context.Context, t *Task) (int64, error)
	UpdateTask(ctx context.Context, t *Task) error
	DeleteTask(ctx context.Context, id int64) error
}
