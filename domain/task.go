package domain

import "time"

type Task struct {
	ID          int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, ErrTitleEmpty
	}
	return &Task{
		Title:       title,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (t *Task) Update(title, description string) error {
	if title == "" {
		return ErrTitleEmpty
	}
	t.Title = title
	t.Description = description
	t.UpdatedAt = time.Now()
	return nil
}
