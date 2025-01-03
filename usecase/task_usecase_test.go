package usecase

import (
	"context"
	"testing"

	"github.com/kazuyainoue0124/go-rest-api/domain"
)

type mockRepo struct {
	tasks  map[int64]*domain.Task
	nextID int64
}

func NewMockTaskRepo() *mockRepo {
	return &mockRepo{
		tasks:  make(map[int64]*domain.Task),
		nextID: 1,
	}
}

func (m *mockRepo) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	var result []*domain.Task
	for _, v := range m.tasks {
		result = append(result, v)
	}
	return result, nil
}

func (m *mockRepo) GetTaskById(ctx context.Context, id int64) (*domain.Task, error) {
	task, ok := m.tasks[id]
	if !ok {
		return nil, domain.ErrNotFound
	}
	return task, nil
}

func (m *mockRepo) CreateTask(ctx context.Context, t *domain.Task) (int64, error) {
	id := m.nextID
	m.nextID++
	t.ID = id
	m.tasks[id] = t
	return id, nil
}

func (m *mockRepo) UpdateTask(ctx context.Context, t *domain.Task) error {
	if _, ok := m.tasks[t.ID]; !ok {
		return domain.ErrNotFound
	}
	m.tasks[t.ID] = t
	return nil
}

func (m *mockRepo) DeleteTask(ctx context.Context, id int64) error {
	if _, ok := m.tasks[id]; !ok {
		return domain.ErrNotFound
	}
	delete(m.tasks, id)
	return nil
}

func TestUsecase_GetAllTasks(t *testing.T) {
	repo := NewMockTaskRepo()
	u := NewTaskUsecase(repo)

	u.CreateTask(context.Background(), "タイトル1", "説明1")
	u.CreateTask(context.Background(), "タイトル2", "説明2")

	tasks, err := u.GetAllTasks(context.Background())
	if err != nil {
		t.Errorf("一覧取得でエラー発生: %v", err)
	}
	if len(tasks) != 2 {
		t.Errorf("タスク数が異なります。期待: 2, 実際: %d", len(tasks))
	}
}

func TestUsecase_GetTaskById(t *testing.T) {
	repo := NewMockTaskRepo()
	u := NewTaskUsecase(repo)

	id, _ := u.CreateTask(context.Background(), "タイトル", "説明")
	task, err := u.GetTaskById(context.Background(), id)
	if err != nil {
		t.Errorf("存在するタスクの取得でエラー発生: %v", err)
	}
	if task.Title != "タイトル" {
		t.Errorf("タイトルが一致しません。期待=%q, 実際=%q", "タイトル", task.Title)
	}

	_, err = u.GetTaskById(context.Background(), 9999)
	if err == nil {
		t.Error("存在しないIDなのでエラーになるはずが、エラーになっていません。")
	}
}

func TestUsecase_CreateTask(t *testing.T) {
	repo := NewMockTaskRepo()
	u := NewTaskUsecase(repo)

	id, err := u.CreateTask(context.Background(), "タイトル", "説明")
	if err != nil {
		t.Errorf("タスク作成でエラー発生: %v", err)
	}
	if id == 0 {
		t.Error("IDが0です。IDは1以上である必要があります")
	}

	_, err = u.CreateTask(context.Background(), "", "説明")
	if err == nil {
		t.Error("タイトルが空なのでエラーになるはずが、エラーになっていません。")
	}
}

func TestUsecase_UpdateTask(t *testing.T) {
	repo := NewMockTaskRepo()
	u := NewTaskUsecase(repo)

	id, _ := u.CreateTask(context.Background(), "タイトル", "説明")
	err := u.UpdateTask(context.Background(), id, "", "更新した説明")
	if err == nil {
		t.Error("タイトルが空なので更新エラーになるはずが出ません。")
	}

	err = u.UpdateTask(context.Background(), id, "更新したタイトル", "更新した説明")
	if err != nil {
		t.Errorf("タスク更新でエラー発生: %v", err)
	}

	err = u.UpdateTask(context.Background(), 9999, "存在しないタイトル", "存在しない説明")
	if err == nil {
		t.Error("存在しないIDなのでエラーになるはずが、エラーになっていません。")
	}
}

func TestUsecase_Delete(t *testing.T) {
	repo := NewMockTaskRepo()
	u := NewTaskUsecase(repo)

	id, _ := u.CreateTask(context.Background(), "タイトル", "説明")
	err := u.DeleteTask(context.Background(), id)
	if err != nil {
		t.Errorf("タスク削除でエラー発生: %v", err)
	}

	err = u.DeleteTask(context.Background(), id)
	if err == nil {
		t.Error("既に削除済みのタスクなのでエラーになるはずが、エラーになっていません。")
	}

	err = u.DeleteTask(context.Background(), 9999)
	if err == nil {
		t.Error("存在しないIDなのでエラーになるはずが、エラーになっていません。")
	}
}
