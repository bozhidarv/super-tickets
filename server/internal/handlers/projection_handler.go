package handlers

import (
	"encoding/json"
	"supertickets/internal/models"
	"supertickets/internal/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func GetProjectionsHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projections, err := repo.ProjectionRepo.GetProjections()
		if err != nil {
			http.Error(w, "Error fetching projections", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(projections)
	}
}


func CreateProjectionHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var proj models.Projection
		if err := json.NewDecoder(r.Body).Decode(&proj); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if err := repo.ProjectionRepo.CreateProjection(&proj); err != nil {
			http.Error(w, "Error creating projection", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(proj)
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
		var proj models.Projection
		if err := json.NewDecoder(r.Body).Decode(&proj); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		proj.ID = id
		if err := repo.ProjectionRepo.UpdateProjection(&proj); err != nil {
			http.Error(w, "Error updating projection", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(proj)
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
