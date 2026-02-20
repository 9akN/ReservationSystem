package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/yourname/reservation-system/internal/models"
	"github.com/yourname/reservation-system/internal/service"
)

type ReservationHandler struct {
	service *service.ReservationService
}

func NewReservationHandler(service *service.ReservationService) *ReservationHandler {
	return &ReservationHandler{service: service}
}

func (h *ReservationHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/reservations", h.handleReservations)
	mux.HandleFunc("/reservations/", h.handleReservationByID)
}

// GET /reservations
// POST /reservations
func (h *ReservationHandler) handleReservations(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.list(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /reservations/:id
// PUT /reservations/:id
// DELETE /reservations/:id
func (h *ReservationHandler) handleReservationByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/reservations/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.getByID(w, r, int32(id))
	case http.MethodPut:
		h.update(w, r, int32(id))
	case http.MethodDelete:
		h.delete(w, r, int32(id))
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *ReservationHandler) list(w http.ResponseWriter, r *http.Request) {
	reservations, err := h.service.List(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, http.StatusOK, reservations)
}

func (h *ReservationHandler) getByID(w http.ResponseWriter, r *http.Request, id int32) {
	reservation, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, http.StatusOK, reservation)
}

func (h *ReservationHandler) create(w http.ResponseWriter, r *http.Request) {
	var req models.CreateReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	reservation, err := h.service.Create(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusCreated, reservation)
}

func (h *ReservationHandler) update(w http.ResponseWriter, r *http.Request, id int32) {
	var req models.UpdateReservationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	reservation, err := h.service.Update(r.Context(), id, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	writeJSON(w, http.StatusOK, reservation)
}

func (h *ReservationHandler) delete(w http.ResponseWriter, r *http.Request, id int32) {
	if err := h.service.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
