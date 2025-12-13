package main

import (
	"log"
	"net/http"
	"os"

	"github.com/MananLedwani/taskmanager/backend/internal/config"
	"github.com/MananLedwani/taskmanager/backend/internal/constants"
	"github.com/MananLedwani/taskmanager/backend/internal/middleware"
	"github.com/MananLedwani/taskmanager/backend/internal/repository"
	"github.com/MananLedwani/taskmanager/backend/internal/router"
	"github.com/MananLedwani/taskmanager/backend/internal/service"
)

func runSQLScript(db *config.Database, filePath string) {
	script, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to read SQL script:", err)
	}

	if _, err := db.DB.Exec(string(script)); err != nil {
		log.Fatal("Failed to execute SQL script:", err)
	}

	log.Println("Database tables created")
}

func main() {

	dbConn, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer dbConn.DB.Close()

	runSQLScript(dbConn, constants.DBUrl)

	userRepo := repository.NewUserRepository(dbConn.DB)
	taskRepo := repository.NewTaskRepository(dbConn.DB)

	authService := service.NewAuthService(userRepo)
	taskService := service.NewTaskService(taskRepo)

	r := router.SetupRouter(authService, taskService)

	handler := middleware.CorsMiddleWare(r)

	log.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("Server failed:", err)
	}
}
