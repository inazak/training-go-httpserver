package httpserver

import (
	"github.com/go-chi/chi/v5"
	"github.com/inazak/training-go-httpserver/httpserver/handler/api"
	"github.com/inazak/training-go-httpserver/service"
	"net/http"
)

func NewMux(svc service.Service) http.Handler {
	mux := chi.NewRouter()

	apiHandler := api.NewHandler(svc)
	mux.HandleFunc("/health", apiHandler.ServeHealthCheck)
	mux.Post("/register", apiHandler.ServeAddUser)
	mux.Post("/login", apiHandler.ServeLogin)

	mux.Route("/task", func(r chi.Router) {
		r.Use(apiHandler.AuthMiddleware)
		r.Get("/", apiHandler.ServeGetTaskList)
		r.Post("/", apiHandler.ServeAddTask)
	})

	return mux
}
