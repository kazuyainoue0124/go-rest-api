package router

import (
	"net/http"

	"github.com/kazuyainoue0124/go-rest-api/presentation/handlers"
)

func NewRouter(h *handlers.TaskHandler) *http.ServeMux {
	// /tasks -> GET: GetAllTasks, POST: CreateTask
	// /tasks/{id} -> GET: GetTaskById, PUT: UpdateTask, DELETE: DeleteTask

	mux := http.NewServeMux()

	// /tasks のパターンを先に登録
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/tasks" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.GetAllTasks(w, r)
		case http.MethodPost:
			h.CreateTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// /tasks/{id} のパターンを登録
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/tasks/" {
			http.NotFound(w, r)
			return
		}

		switch r.Method {
		case http.MethodGet:
			h.GetTaskById(w, r)
		case http.MethodPut:
			h.UpdateTask(w, r)
		case http.MethodDelete:
			h.DeleteTask(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
