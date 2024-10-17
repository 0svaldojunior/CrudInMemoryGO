package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)

	router.Post("/api/users", CreateUser)
	router.Get("/api/users", GetAllUsers)
	router.Get("/api/users/{id}", GetUser)
	router.Put("/api/users/{id}", PutUser)
	router.Delete("/api/users/{id}", DeleteUser)

	return router
}
