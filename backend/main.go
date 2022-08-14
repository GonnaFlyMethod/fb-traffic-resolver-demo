package main

import (
	"backend/server/handler"
	"backend/storage"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	log.Println("starting demo server...")

	inMemoryStorage := storage.NewInMemoryStorage()
	appHandler := handler.NewHandler(inMemoryStorage)

	router := chi.NewRouter()
	appHandler.BuildAccessPoliciesAndRoutes(router)

	server := &http.Server{Addr: ":8000", Handler: router}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("can't start demo server")
	}
}
