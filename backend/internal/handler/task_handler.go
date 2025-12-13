package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MananLedwani/taskmanager/backend/internal/middleware"
	"github.com/MananLedwani/taskmanager/backend/internal/models"
	"github.com/MananLedwani/taskmanager/backend/internal/service"
	"github.com/google/uuid"
)

type TaskHandler struct {
	taskService service.TaskService
}

func NewTaskHandler(taskService service.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: taskService,
	}
}

func getUserID(r *http.Request) (uuid.UUID, error) {
	userIDStr := r.Context().Value(middleware.UserIDKey).(string)
	return uuid.Parse(userIDStr)
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, "Invalid user", http.StatusUnauthorized)
		return
	}

	tasks, err := h.taskService.GetTasksByUser(userID)
	if err != nil {
		http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task.UserID = userID

	if err := h.taskService.CreateTask(&task); err != nil {
		http.Error(w, "Failed to create task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	taskIDStr := r.PathValue("id")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	task.ID = taskID
	task.UserID = userID

	if err := h.taskService.UpdateTask(&task); err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task updated successfully",
	})
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserID(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	taskIDStr := r.PathValue("id")
	taskID, err := uuid.Parse(taskIDStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	if err := h.taskService.DeleteTask(taskID, userID); err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Task deleted successfully",
	})
}
