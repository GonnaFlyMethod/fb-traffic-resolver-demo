package handler

import (
	"bytes"
	"encoding/json"
	"github.com/GonnaFlyMethod/fb-traffic-resolver-demo/backend/common"
	"github.com/GonnaFlyMethod/fb-traffic-resolver-demo/backend/server"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
)

type usersListResponse struct {
	Users []*common.UserWithID `json:"users"`
	Count int                  `json:"count"`
}

type UserStorage interface {
	GetAllUsers() []*common.UserWithID
	CreateNewUser(u common.UserWithoutID) error
	UpdateUser(userID string, user common.UserWithoutID) bool
	DeleteUser(userID string) bool
}

type Handler struct {
	storage UserStorage
}

func NewHandler(userStorage UserStorage) *Handler {
	return &Handler{storage: userStorage}
}

func (h *Handler) BuildAccessPoliciesAndRoutes(router *chi.Mux) {
	router.Use(server.ResponseAsJSONMiddleware)
	router.Use(server.CorsMiddleware)

	router.Get("/api/users", h.getAllUsers)
	router.Post("/api/users", h.createNewUser)
	router.Put("/api/users/{id}", h.updateUser)
	router.Delete("/api/users/{id}", h.deleteUser)
}

func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users := h.storage.GetAllUsers()
	responseBody := usersListResponse{Users: users, Count: len(users)}
	writeResponse(w, http.StatusOK, responseBody)
}

func (h *Handler) createNewUser(w http.ResponseWriter, r *http.Request) {
	requestBodyAsBytes, err := readRequestBody(r)
	if err != nil {
		writeInternalErrResponse(w)
		return
	}

	jsonDecoder := json.NewDecoder(bytes.NewReader(requestBodyAsBytes))
	var user common.UserWithoutID
	if err = jsonDecoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !user.IsCompletelyFilled() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.storage.CreateNewUser(user); err != nil {
		writeInternalErrResponse(w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	userIDToUpdate := chi.URLParam(r, "id")
	if _, err := uuid.Parse(userIDToUpdate); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestBodyAsBytes, err := readRequestBody(r)
	if err != nil {
		writeInternalErrResponse(w)
		return
	}

	jsonDecoder := json.NewDecoder(bytes.NewReader(requestBodyAsBytes))
	var user common.UserWithoutID
	if err = jsonDecoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !user.IsCompletelyFilled() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isUserUpdated := h.storage.UpdateUser(userIDToUpdate, user)
	if !isUserUpdated {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	userIDToDelete := chi.URLParam(r, "id")
	if _, err := uuid.Parse(userIDToDelete); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	isUserDeleted := h.storage.DeleteUser(userIDToDelete)
	if !isUserDeleted {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
