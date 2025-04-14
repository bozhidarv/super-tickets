package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"supertickets/internal/auth"
	"supertickets/internal/models"
	"supertickets/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// RegisterHandler handles user registration and returns a JWT token in the Authorization header.
func RegisterHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
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
		user.Role = "user" // default role.

		// Save the user.
		if err := repo.UserRepo.CreateUser(&user); err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		// Generate a JWT token for the newly registered user.
		token, err := auth.GenerateToken(&user)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Set the token in the Authorization header.
		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusCreated)

		// Optionally, return user info in the JSON response (without the token).
		json.NewEncoder(w).Encode(user)
	}
}

// LoginHandler handles user authentication and returns a JWT token in the Authorization header.
func LoginHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		slog.Debug(creds.Username)
		// Retrieve the user by username.
		user, err := repo.UserRepo.GetUserByUsername(creds.Username)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Compare the hashed password.
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Generate a JWT token for the authenticated user.
		token, err := auth.GenerateToken(user)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		// Set the token in the Authorization header.
		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusOK)

		// Optionally, return a success message.
		json.NewEncoder(w).Encode(map[string]string{"message": "login successful"})
	}
}
