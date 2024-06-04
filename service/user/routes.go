package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafael-men/rest-api-with-golang/service/auth"
	"github.com/rafael-men/rest-api-with-golang/types"
	"github.com/rafael-men/rest-api-with-golang/utils"
)

type Handler struct {
	Store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// Implementação do login
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	err := utils.ParseJSON(r, &payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// check if the user exists
	_, err = h.Store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with that email already exists: %s", payload.Email))
		return
	}

	hashedPassword, err := auth.HashPassword(payload.Password)

	err = h.Store.CreateUser(types.User{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		Password:  hashedPassword,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)

}
