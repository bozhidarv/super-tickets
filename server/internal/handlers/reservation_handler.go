package handlers

import (
	"encoding/json"
	"moviereservationsystem/internal/auth"
	"moviereservationsystem/internal/middleware"
	"moviereservationsystem/internal/models"
	"moviereservationsystem/internal/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateReservationHandler creates a new reservation.
func CreateReservationHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resv models.Reservation
		if err := json.NewDecoder(r.Body).Decode(&resv); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		claims, ok := r.Context().Value(middleware.ContextUserKey).(*auth.Claims)
		if !ok || claims == nil {
			http.Error(w, "Unable to retrieve user info from token", http.StatusUnauthorized)
			return
		}

		resv.UserID = claims.UserID
		if err := repo.ReservationRepo.CreateReservation(&resv); err != nil {
			http.Error(w, "Error creating reservation", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(resv)
	}
}

// GetReservationsHandler returns reservations for a given user.
// (For simplicity, the user ID is taken from a query parameter. In production, you would extract it from the JWT.)
func GetReservationsHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the claims from the context (set by the auth middleware)
		claims, ok := r.Context().Value(middleware.ContextUserKey).(*auth.Claims)
		if !ok || claims == nil {
			http.Error(w, "Unable to retrieve user info from token", http.StatusUnauthorized)
			return
		}

		// Use the user ID from the JWT token
		userID := claims.UserID

		reservations, err := repo.ReservationRepo.GetReservationsByUser(userID)
		if err != nil {
			http.Error(w, "Error fetching reservations", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(reservations)
	}
}

// DeleteReservationHandler cancels a reservation.
func DeleteReservationHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
			return
		}
		if err := repo.ReservationRepo.DeleteReservation(id); err != nil {
			http.Error(w, "Error deleting reservation", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
