package handlers

import (
	"encoding/json"
	"moviereservationsystem/internal/models"
	"moviereservationsystem/internal/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// GetUsersHandler returns a list of all users (admin-only).
func GetUsersHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := repo.UserRepo.GetAllUsers()
		if err != nil {
			http.Error(w, "Error fetching users", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

// CreateUserHandler allows an admin to create a new user.
func CreateUserHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Validate required fields.
		if user.Username == "" || user.Password == "" || user.Role == "" {
			http.Error(w, "Username, password, and role are required", http.StatusBadRequest)
			return
		}

		// Hash the password.
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)

		// Save the user.
		if err := repo.UserRepo.CreateUser(&user); err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}

// UpdateUserHandler allows an admin to update an existing user's details.
func UpdateUserHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the user ID from the URL.
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		user.ID = id

		// If a new password is provided, hash it; otherwise, preserve the existing password.
		if user.Password != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword(
				[]byte(user.Password),
				bcrypt.DefaultCost,
			)
			if err != nil {
				http.Error(w, "Error hashing password", http.StatusInternalServerError)
				return
			}
			user.Password = string(hashedPassword)
		} else {
			// Retrieve the existing user to preserve the current password.
			existingUser, err := repo.UserRepo.GetUserById(id)
			if err != nil {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			user.Password = existingUser.Password
		}

		// Update the user.
		if err := repo.UserRepo.UpdateUser(&user); err != nil {
			http.Error(w, "Error updating user", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(user)
	}
}

// DeleteUserHandler allows an admin to delete a user.
func DeleteUserHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		if err := repo.UserRepo.DeleteUser(id); err != nil {
			http.Error(w, "Error deleting user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
