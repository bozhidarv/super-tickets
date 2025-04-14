package server

import (
	"net/http"
	"supertickets/internal/handlers"
	"supertickets/internal/middleware"
	"supertickets/internal/repository"

	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
	Repo   *repository.Repository
}

func NewServer(repo *repository.Repository) *Server {
	s := &Server{
		Router: mux.NewRouter(),
		Repo:   repo,
	}

	s.Router.Use(middleware.CORSMiddleware)
	s.routes()
	return s
}

func (s *Server) routes() {
	s.Router.PathPrefix("/").
		Methods(http.MethodOptions).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

	api := s.Router.PathPrefix("/api").Subrouter()

	api.HandleFunc("/auth/register", handlers.RegisterHandler(s.Repo)).Methods("POST")
	api.HandleFunc("/auth/login", handlers.LoginHandler(s.Repo)).Methods("POST")

	protected := api.PathPrefix("").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	protected.HandleFunc("/movies", handlers.GetMoviesHandler(s.Repo)).Methods("GET")
	protected.HandleFunc("/movies", handlers.CreateMovieHandler(s.Repo)).Methods("POST")
	protected.HandleFunc("/movies/{id}", handlers.UpdateMovieHandler(s.Repo)).Methods("PUT")
	protected.HandleFunc("/movies/{id}", handlers.DeleteMovieHandler(s.Repo)).Methods("DELETE")

	protected.HandleFunc("/projections", handlers.GetProjectionsHandler(s.Repo)).Methods("GET")
	protected.HandleFunc("/projections", handlers.CreateProjectionHandler(s.Repo)).Methods("POST")
	protected.HandleFunc("/projections/{id}", handlers.UpdateProjectionHandler(s.Repo)).
		Methods("PUT")
	protected.HandleFunc("/projections/{id}", handlers.DeleteProjectionHandler(s.Repo)).
		Methods("DELETE")

	protected.HandleFunc("/reservations", handlers.CreateReservationHandler(s.Repo)).Methods("POST")
	protected.HandleFunc("/reservations", handlers.GetReservationsHandler(s.Repo)).Methods("GET")
	protected.HandleFunc("/reservations/{id}", handlers.DeleteReservationHandler(s.Repo)).
		Methods("DELETE")

	protected.HandleFunc("/users", handlers.GetUsersHandler(s.Repo)).Methods("GET")
	protected.HandleFunc("/users", handlers.CreateUserHandler(s.Repo)).Methods("POST")
	protected.HandleFunc("/users/{id}", handlers.UpdateUserHandler(s.Repo)).Methods("PUT")
	protected.HandleFunc("/users/{id}", handlers.DeleteUserHandler(s.Repo)).Methods("DELETE")
}
