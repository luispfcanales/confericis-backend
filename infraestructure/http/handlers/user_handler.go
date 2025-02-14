package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/confericis-backend/model"
	"github.com/confericis-backend/ports/input"
)

type UserHandler struct {
	userService input.UserService
}

func NewUserHandler(us input.UserService) *UserHandler {
	return &UserHandler{
		userService: us,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		RoleID   string `json:"role_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &model.User{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
		RoleID:   req.RoleID,
	}

	if err := h.userService.CreateUser(r.Context(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
		"name":  user.Name,
	})
}
