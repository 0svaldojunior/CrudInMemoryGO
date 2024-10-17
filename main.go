package main

import (
	"CrudInMemoryGo/api"
	"CrudInMemoryGo/log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Error("Failed to execute code", err)
		os.Exit(1)
	}

	log.Info("Server started successfully")
}

func run() error {
	log.Info("Starting server‼ 🚀")

	handler := api.NewHandler()

	server := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
