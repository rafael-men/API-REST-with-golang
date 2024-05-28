package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafael-men/rest-api-with-golang/types"
	"github.com/rafael-men/rest-api-with-golang/utils"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
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

	// err := userService.RegisterUser(payload)
	// if err != nil {
	//     http.Error(w, "Error while registering user", http.StatusInternalServerError)
	//     return
	// }

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Successfully Registered"))

}
