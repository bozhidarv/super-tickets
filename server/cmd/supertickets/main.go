package main

import (
	"log"
	"net/http"
	"supertickets/internal/repository"
	"supertickets/internal/server"
	"supertickets/internal/utils"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	utils.LoadEnvVars()

	// Initialize the PostgreSQL connection using SQLX
	db, err := repository.NewPostgresDB(utils.EnvVars.DbUrl())
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
		Addr:         ":" + utils.EnvVars.Port(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting on :8080")
	log.Fatal(httpServer.ListenAndServe())
}
