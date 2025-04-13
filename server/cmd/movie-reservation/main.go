package main

import (
	"log"
	"moviereservationsystem/internal/repository"
	"moviereservationsystem/internal/server"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	// Read configuration (for simplicity, using an environment variable with a default)
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/moviereservation?sslmode=disable"
	}

	// Initialize the PostgreSQL connection using SQLX
	db, err := repository.NewPostgresDB(dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create the composite repository which bundles all model-specific repos.
	repo := repository.NewRepository(db)

	// Create the server and register the routes.
	srv := server.NewServer(repo)

	// Create and configure the HTTP server.
	httpServer := &http.Server{
		Handler:      srv.Router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting on :8080")
	log.Fatal(httpServer.ListenAndServe())
}
