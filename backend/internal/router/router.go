package router

import (
	"net/http"

	"github.com/MananLedwani/taskmanager/backend/internal/handler"
	"github.com/MananLedwani/taskmanager/backend/internal/middleware"
	"github.com/MananLedwani/taskmanager/backend/internal/service"
)

func SetupRouter(
	authService service.AuthService,
	taskService service.TaskService,
) *http.ServeMux {

	r := http.NewServeMux()

	authHandler := handler.NewAuthHandler(authService)
	taskHandler := handler.NewTaskHandler(taskService)

	r.HandleFunc("POST /signup", authHandler.Signup)
	r.HandleFunc("POST /login", authHandler.Login)

	r.Handle(
		"GET /tasks",
		middleware.JWTAuthMiddleware(http.HandlerFunc(taskHandler.GetTasks)),
	)

	r.Handle(
		"POST /tasks",
		middleware.JWTAuthMiddleware(http.HandlerFunc(taskHandler.CreateTask)),
	)

	r.Handle(
		"PATCH /tasks/{id}",
		middleware.JWTAuthMiddleware(http.HandlerFunc(taskHandler.UpdateTask)),
	)

	r.Handle(
		"DELETE /tasks/{id}",
		middleware.JWTAuthMiddleware(http.HandlerFunc(taskHandler.DeleteTask)),
	)

	return r
}
