package handler

import (
	"backend/common"
	"backend/server"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

type usersListResponse struct {
	Users []common.UserWithID `json:"users"`
	Count int                 `json:"count"`
}

type UserStorage interface {
	GetAllUsers() []common.UserWithID
	CreateNewUser(u common.UserWithoutID) error
	// UpdateUser()
	// DeleteUser()
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

	router.Get("/users", h.getAllUsers)
	router.Post("/users", h.createNewUser)
	router.Put("/users/{id}", h.updateUser)
	router.Delete("/users/{id}", h.deleteUser)
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
	fmt.Println(r, "update")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}

func readRequestBody(r *http.Request) ([]byte, error) {
	content, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	defer func(r *http.Request) {
		if err = r.Body.Close(); err != nil {
			fmt.Println("failed to close request body")
		}
	}(r)

	r.Body = io.NopCloser(bytes.NewBuffer(content))
	return content, nil
}
