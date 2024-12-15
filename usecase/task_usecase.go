package usecase

import (
	"context"

	"github.com/kazuyainoue0124/go-rest-api/domain"
)

type TaskUsecase struct {
	repo domain.ITaskRepository
}

func NewTaskUsecase(repo domain.ITaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

// 全件取得
func (u *TaskUsecase) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	return u.repo.GetAllTasks(ctx)
}

// 1件取得
func (u *TaskUsecase) GetTaskById(ctx context.Context, id int64) (*domain.Task, error) {
	return u.repo.GetTaskById(ctx, id)
}

// 作成
func (u *TaskUsecase) CreateTask(ctx context.Context, title, description string) (int64, error) {
	if title == "" {
		return 0, domain.ErrInvalid
	}
	t := &domain.Task{
		Title:       title,
		Description: description,
	}

	return u.repo.CreateTask(ctx, t)
}

// 更新
func (u *TaskUsecase) UpdateTask(ctx context.Context, id int64, title, description string) error {
	t, err := u.repo.GetTaskById(ctx, id)
	if err != nil {
		return err
	}
	if title == "" {
		return domain.ErrInvalid
	}
	t.Title = title
	t.Description = description
	return u.repo.UpdateTask(ctx, t)
}

// 削除
func (u *TaskUsecase) DeleteTask(ctx context.Context, id int64) error {
	return u.repo.DeleteTask(ctx, id)
}
