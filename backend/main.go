package main

import (
	"fmt"
	"github.com/GonnaFlyMethod/fb-traffic-resolver-demo/backend/server/handler"
	"github.com/GonnaFlyMethod/fb-traffic-resolver-demo/backend/storage"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
)

func mustGetEnv(envName string) string {
	res := os.Getenv(envName)
	if res == "" {
		panicMsg := fmt.Sprintf("can't read env %s", envName)
		panic(panicMsg)
	}
	return res
}

func main() {
	log.Println("starting demo server...")

	inMemoryStorage := storage.NewInMemoryStorage()
	appHandler := handler.NewHandler(inMemoryStorage)

	router := chi.NewRouter()
	appHandler.BuildAccessPoliciesAndRoutes(router)

	serverPort := mustGetEnv("BACKEND_PORT")
	addr := fmt.Sprintf(":%s", serverPort)

	server := &http.Server{Addr: addr, Handler: router}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("can't start demo server")
	}
}
