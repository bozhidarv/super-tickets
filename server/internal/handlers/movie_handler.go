package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"supertickets/internal/models"
	"supertickets/internal/repository"

	"github.com/gorilla/mux"
)

func GetMoviesHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movies, err := repo.MovieRepo.GetMovies()
		if err != nil {
			http.Error(w, "Error fetching movies", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(movies)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
	}
}

func CreateMovieHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie models.Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if err := repo.MovieRepo.CreateMovie(&movie); err != nil {
			http.Error(w, "Error creating movie", http.StatusInternalServerError)
			return
		}

		err := json.NewEncoder(w).Encode(movie)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateMovieHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}
		var movie models.Movie
		if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		movie.ID = id
		if err := repo.MovieRepo.UpdateMovie(&movie); err != nil {
			http.Error(w, "Error updating movie", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(movie)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}

func DeleteMovieHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid movie ID", http.StatusBadRequest)
			return
		}
		if err := repo.MovieRepo.DeleteMovie(id); err != nil {
			http.Error(w, "Error deleting movie", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
