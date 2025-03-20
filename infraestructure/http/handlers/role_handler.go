package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luispfcanales/confericis-backend/ports/input"
)

type RoleHandler struct {
	roleService input.RoleService
}

func NewRoleHandler(rs input.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: rs,
	}
}

// obtine todos los roles
func (h *RoleHandler) AllRoles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	roles, err := h.roleService.GetRoles(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(roles)
}
