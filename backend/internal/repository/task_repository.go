package repository

import (
	"database/sql"

	"github.com/MananLedwani/taskmanager/backend/internal/models"
	"github.com/google/uuid"
)

type TaskRepository interface {
	Create(task *models.Task) error
	GetAllByUser(userID uuid.UUID) ([]models.Task, error)
	Update(task *models.Task) error
	Delete(taskID uuid.UUID, userID uuid.UUID) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.Task) error {

	query := `
		INSERT INTO tasks (title, user_id)
		VALUES ($1, $2)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		task.Title,
		task.UserID,
	).Scan(&task.ID)
}

func (r *taskRepository) GetAllByUser(userID uuid.UUID) ([]models.Task, error) {

	query := `
		SELECT id, title, user_id
		FROM tasks
		WHERE user_id = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.UserID,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *taskRepository) Update(task *models.Task) error {

	query := `
		UPDATE tasks
		SET title = $1
		WHERE id = $2 AND user_id = $3
	`

	_, err := r.db.Exec(
		query,
		task.Title,
		task.ID,
		task.UserID,
	)

	return err
}

func (r *taskRepository) Delete(taskID uuid.UUID, userID uuid.UUID) error {

	query := `
		DELETE FROM tasks
		WHERE id = $1 AND user_id = $2
	`

	_, err := r.db.Exec(query, taskID, userID)
	return err
}
