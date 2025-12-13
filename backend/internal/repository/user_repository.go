package repository

import (
	"database/sql"
	"errors"

	"github.com/MananLedwani/taskmanager/backend/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {

	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := r.db.QueryRow(
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(&user.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {

	query := `
		SELECT id, username, email, password
		FROM users
		WHERE email = $1
	`

	user := &models.User{}

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}
