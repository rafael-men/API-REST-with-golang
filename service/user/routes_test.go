package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rafael-men/rest-api-with-golang/types"
)

func TestUserServiceHandlers(t *testing.T) {

	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("fail if user payload is invalid", func(t *testing.T) {
		payload := types.RegisterUserPayload{
			Firstname: "user",
			Lastname:  "123",
			Email:     "",
			Password:  "asdfghjklm",
		}
		marshalled, _ := json.Marshal(payload)
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister).Methods("POST")
		router.ServeHTTP(rr, req)

		// Verifica o status de resposta e a mensagem de erro esperada
		if rr.Code != http.StatusBadRequest {
			t.Errorf("Status Code Expected, got %d", http.StatusBadRequest, rr.Code)
		}

		expected := `{"error":"Key: 'RegisterUserPayload.Email' Error:Field validation for 'Email' failed on the 'required' tag"}`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}
