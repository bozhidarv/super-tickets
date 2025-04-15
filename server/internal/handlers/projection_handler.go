package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"supertickets/internal/models"
	"supertickets/internal/repository"

	"github.com/gorilla/mux"
)

func GetProjectionsHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projections, err := repo.ProjectionRepo.GetProjections()
		if err != nil {
			http.Error(w, "Error fetching projections", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(projections)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
	}
}

func CreateProjectionHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var projection models.Projection
		if err := json.NewDecoder(r.Body).Decode(&projection); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if err := repo.ProjectionRepo.CreateProjection(&projection); err != nil {
			http.Error(w, "Error creating projection", http.StatusInternalServerError)
			return
		}

		err := json.NewEncoder(w).Encode(projection)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateProjectionHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid projection ID", http.StatusBadRequest)
			return
		}
		var projection models.Projection
		if err := json.NewDecoder(r.Body).Decode(&projection); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		projection.ID = id
		if err := repo.ProjectionRepo.UpdateProjection(&projection); err != nil {
			http.Error(w, "Error updating projection", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(projection)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
	}
}

func DeleteProjectionHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid projection ID", http.StatusBadRequest)
			return
		}
		if err := repo.ProjectionRepo.DeleteProjection(id); err != nil {
			http.Error(w, "Error deleting projection", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
