package api

import (
	"encoding/json"
	"net/http"

	"github.com/yourorg/auth-service/internal/models"
	"github.com/yourorg/auth-service/internal/service"
)

type Handler struct {
	validator *service.Validator
}

func NewHandler(validator *service.Validator) *Handler {
	return &Handler{
		validator: validator,
	}
}

func (h *Handler) ValidateAccess(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.AccessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.validator.ValidateAccess(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
