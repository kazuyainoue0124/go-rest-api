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
	task, err := domain.NewTask(title, description)
	if err != nil {
		return 0, err
	}

	id, err := u.repo.CreateTask(ctx, task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// 更新
func (u *TaskUsecase) UpdateTask(ctx context.Context, id int64, title, description string) error {
	task, err := u.repo.GetTaskById(ctx, id)
	if err != nil {
		return err
	}
	if err := task.Update(title, description); err != nil {
		return err
	}

	if err := u.repo.UpdateTask(ctx, task); err != nil {
		return err
	}
	return nil
}

// 削除
func (u *TaskUsecase) DeleteTask(ctx context.Context, id int64) error {
	return u.repo.DeleteTask(ctx, id)
}
