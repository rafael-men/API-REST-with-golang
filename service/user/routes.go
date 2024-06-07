package user

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rafael-men/rest-api-with-golang/types"
)

type UserStore interface {
	GetUserByEmail(email string) (*types.User, error)
	GetUserByID(id int) (*types.User, error)
	CreateUser(user types.User) error
}

type Handler struct {
	userStore UserStore
	validator *validator.Validate
}

func NewHandler(store UserStore) *Handler {
	return &Handler{
		userStore: store,
		validator: validator.New(),
	}
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var payload types.RegisterUserPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perform input validation here
	if err := h.validator.Struct(payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Check if user already exists
	if _, err := h.userStore.GetUserByEmail(payload.Email); err == nil {
		http.Error(w, "user with that email already exists", http.StatusConflict)
		return
	}

	// Create the user
	newUser := types.User{
		Firstname: payload.Firstname,
		Lastname:  payload.Lastname,
		Email:     payload.Email,
		Password:  payload.Password, // Make sure to hash the password in real implementations
	}

	if err := h.userStore.CreateUser(newUser); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
