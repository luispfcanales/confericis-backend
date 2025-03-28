package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luispfcanales/confericis-backend/ports/input"
)

type DriveHandler struct {
	driveService input.DriveService
}

func NewDriveHandler(driveService input.DriveService) *DriveHandler {
	return &DriveHandler{
		driveService: driveService,
	}
}

func (h *DriveHandler) ListFiles(w http.ResponseWriter, r *http.Request) {
	parentID := r.PathValue("id")
	if parentID == "" {
		http.Error(w, "folder ID is required", http.StatusBadRequest)
		return
	}

	files, err := h.driveService.ListFiles(r.Context(), parentID)
	if err != nil {
		http.Error(w, "failed to list files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}

func (h *DriveHandler) ListDir(w http.ResponseWriter, r *http.Request) {
	parentID := r.PathValue("id")
	files, err := h.driveService.ListFolders(r.Context(), parentID)
	if err != nil {
		http.Error(w, "failed to list files: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(files)
}
