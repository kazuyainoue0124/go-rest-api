package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kazuyainoue0124/go-rest-api/domain"
	"github.com/kazuyainoue0124/go-rest-api/usecase"
)

type TaskHandler struct {
	u *usecase.TaskUsecase
}

func NewTaskHandler(u *usecase.TaskUsecase) *TaskHandler {
	return &TaskHandler{u: u}
}

// GetAllTasks: GET /tasks
func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.u.GetAllTasks(context.Background())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(tasks)
}

// GetTaskById: GET /tasks/{id}
func (h *TaskHandler) GetTaskById(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("GetTaskById called with path: %s\n", r.URL.Path) // デバッグログ追加

	id, err := extractIdFromPath(r.URL.Path, "/tasks/")
	if err != nil {
		fmt.Printf("Error extracting ID: %v\n", err) // デバッグログ追加
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	fmt.Printf("Extracted ID: %d\n", id) // デバッグログ追加

	task, err := h.u.GetTaskById(context.Background(), id)
	if err != nil {
		fmt.Printf("Error getting task: %v\n", err) // デバッグログ追加
		if errors.Is(err, domain.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(task)
}

// CreateTask: POST /tasks
func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	id, err := h.u.CreateTask(context.Background(), req.Title, req.Description)
	if err != nil {
		if errors.Is(err, domain.ErrInvalid) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// UpdateTask: PUT /tasks/{id}
func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/tasks/")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var req UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err = h.u.UpdateTask(context.Background(), id, req.Title, req.Description)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		if errors.Is(err, domain.ErrInvalid) {
			http.Error(w, "invalid input", http.StatusBadRequest)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

// DeleteTask: DELETE /tasks/{id}
func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := extractIdFromPath(r.URL.Path, "/tasks/")
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	err = h.u.DeleteTask(context.Background(), id)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}

// extractIdFromPath: "/tasks/{id}"からidを抽出
func extractIdFromPath(path, prefix string) (int64, error) {
	if !strings.HasPrefix(path, prefix) {
		return 0, errors.New("invalid path")
	}
	idStr := strings.TrimPrefix(path, prefix)
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}
