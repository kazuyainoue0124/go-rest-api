package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/kazuyainoue0124/go-rest-api/domain"
)

type MySQLTaskRepository struct {
	db *sql.DB
}

func NewMySQLTaskRepository(db *sql.DB) *MySQLTaskRepository {
	return &MySQLTaskRepository{db: db}
}

func (r *MySQLTaskRepository) GetAllTasks(ctx context.Context) ([]*domain.Task, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT id, title, description, created_at, updated_at FROM tasks ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		var t domain.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	return tasks, rows.Err()
}

func (r *MySQLTaskRepository) GetTaskById(ctx context.Context, id int64) (*domain.Task, error) {
	row := r.db.QueryRowContext(ctx,
		"SELECT id, title, description, created_at, updated_at FROM tasks WHERE id = ?", id)
	var t domain.Task
	if err := row.Scan(&t.ID, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &t, nil
}

func (r *MySQLTaskRepository) CreateTask(ctx context.Context, t *domain.Task) (int64, error) {
	stmt, err := r.db.PrepareContext(ctx,
		"INSERT INTO tasks (title, description) VALUES (?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.ExecContext(ctx, t.Title, t.Description)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *MySQLTaskRepository) UpdateTask(ctx context.Context, t *domain.Task) error {
	stmt, err := r.db.PrepareContext(ctx,
		"UPDATE tasks SET title = ?, description = ? WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, t.Title, t.Description, t.ID)
	return err
}

func (r *MySQLTaskRepository) DeleteTask(ctx context.Context, id int64) error {
	stmt, err := r.db.PrepareContext(ctx, "DELETE FROM tasks WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, id)
	return err
}
