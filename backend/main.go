package main

import (
	"backend/server"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	log.Println("starting demo server...")

	h := server.Handler{}
	r := chi.NewRouter()
	h.BuildAccessPoliciesAndRoutes(r)

	s := &http.Server{Addr: ":8000", Handler: r}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("can't start demo server")
	}
}
