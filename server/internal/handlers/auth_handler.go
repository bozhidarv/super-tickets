package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"supertickets/internal/auth"
	"supertickets/internal/models"
	"supertickets/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(repo *repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, &user); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		user.Password = string(hashedPassword)
		user.Role = "user"

		if err := repo.UserRepo.CreateUser(&user); err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		token, err := auth.GenerateToken(&user)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}

		w.Header().Set("Authorization", "Bearer "+token)
		w.WriteHeader(http.StatusCreated)
	}
}

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

		user, err := repo.UserRepo.GetUserByUsername(creds.Username)
		if err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		token, err := auth.GenerateToken(user)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}

		w.Header().Set("Authorization", "Bearer "+token)
	}
}
