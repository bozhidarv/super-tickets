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

	
	db, err := repository.NewPostgresDB(utils.EnvVars.DbUrl())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	
	repo := repository.NewRepository(db)

	
	srv := server.NewServer(repo)

	
	httpServer := &http.Server{
		Handler:      srv.Router,
		Addr:         ":" + utils.EnvVars.Port(),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server starting on :8080")
	log.Fatal(httpServer.ListenAndServe())
}
