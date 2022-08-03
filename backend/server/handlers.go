package server

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type usersListResponse struct {
	Users []map[string]string `json:"users"`
	Count int                 `json:"count"`
}

type Handler struct {
}

func (h *Handler) BuildAccessPoliciesAndRoutes(router *chi.Mux) {
	router.Use(responseAsJSONMiddleware)
	router.Use(corsMiddleware)

	router.Get("/users", h.getAllUsers)
	router.Post("/users", h.createNewUser)
	router.Put("/users/{id}", h.updateUser)
	router.Delete("/users/{id}", h.deleteUser)
}

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []map[string]string{
		{"name": "Test1", "surname": "Test1"},
		{"name": "Test2", "surname": "Test2"},
		{"name": "Test3", "surname": "Test3"},
	}

	responseBody := usersListResponse{
		Users: users,
		Count: 3,
	}

	marshalledResponseBody, err := json.Marshal(responseBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err = w.Write(marshalledResponseBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) createNewUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
