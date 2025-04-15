package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"supertickets/internal/auth"
	"supertickets/internal/middleware"
	"supertickets/internal/models"
	"supertickets/internal/repository"

	"github.com/gorilla/mux"
)

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

		err := json.NewEncoder(w).Encode(resv)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetReservationsHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(middleware.ContextUserKey).(*auth.Claims)
		if !ok || claims == nil {
			http.Error(w, "Unable to retrieve user info from token", http.StatusUnauthorized)
			return
		}

		userID := claims.UserID

		reservations, err := repo.ReservationRepo.GetReservationsByUser(userID)
		if err != nil {
			http.Error(w, "Error fetching reservations", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(reservations)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
	}
}

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
