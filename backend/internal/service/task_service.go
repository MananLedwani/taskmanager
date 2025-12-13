package service

import (
	"github.com/MananLedwani/taskmanager/backend/internal/models"
	"github.com/MananLedwani/taskmanager/backend/internal/repository"
	"github.com/google/uuid"
)

type TaskService interface {
	CreateTask(task *models.Task) error
	GetTasksByUser(userID uuid.UUID) ([]models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(taskID uuid.UUID, userID uuid.UUID) error
}

type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(taskRepo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: taskRepo,
	}
}

func (s *taskService) CreateTask(task *models.Task) error {
	return s.taskRepo.Create(task)
}

func (s *taskService) GetTasksByUser(userID uuid.UUID) ([]models.Task, error) {
	return s.taskRepo.GetAllByUser(userID)
}

func (s *taskService) UpdateTask(task *models.Task) error {
	return s.taskRepo.Update(task)
}

func (s *taskService) DeleteTask(taskID uuid.UUID, userID uuid.UUID) error {
	return s.taskRepo.Delete(taskID, userID)
}
