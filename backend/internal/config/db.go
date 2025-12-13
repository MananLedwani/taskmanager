package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MananLedwani/taskmanager/backend/internal/constants"
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

func NewPostgresDB() (*Database, error) {

	host := constants.DB_HOST
	port := constants.DB_PORT
	user := constants.DB_USER
	password := constants.DB_PASSWORD
	dbname := constants.DB_NAME

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Connected to PostgreSQL")

	return &Database{DB: db}, nil
}