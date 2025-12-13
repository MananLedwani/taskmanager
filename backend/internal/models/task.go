package models

import "github.com/google/uuid"

type Task struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	UserID uuid.UUID `json:"user_id"`
}
