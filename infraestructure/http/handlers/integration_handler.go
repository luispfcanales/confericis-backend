package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/luispfcanales/confericis-backend/model"
)

type IntegrationHandler struct {
	otiURL string
	daaURL string
}

func NewIntegrationHandler(otiURL, daaURL string) *IntegrationHandler {
	return &IntegrationHandler{
		otiURL: otiURL,
		daaURL: daaURL,
	}
}

func (h *IntegrationHandler) fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	return nil
}

func (h *IntegrationHandler) GetReniecInfo(w http.ResponseWriter, r *http.Request) {
	dni := r.PathValue("dni")
	if dni == "" {
		http.Error(w, "DNI parameter is required", http.StatusBadRequest)
		return
	}

	var person model.ReniecPerson
	url := fmt.Sprintf("%s/%s", h.otiURL, dni)
	if err := h.fetchData(url, &person); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}

func (h *IntegrationHandler) GetStudentInfo(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		http.Error(w, "Student code parameter is required", http.StatusBadRequest)
		return
	}

	var student model.UnamadStudent
	url := fmt.Sprintf("%s/student/%s", h.daaURL, code)
	if err := h.fetchData(url, &student); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (h *IntegrationHandler) GetTeacherInfo(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	if code == "" {
		http.Error(w, "Student code parameter is required", http.StatusBadRequest)
		return
	}

	var teacher model.UnamadTeacher
	url := fmt.Sprintf("%s/student/%s", h.daaURL, code)
	if err := h.fetchData(url, &teacher); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teacher)
}
