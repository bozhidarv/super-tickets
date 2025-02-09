package middleware

import (
	"net/http"
)

// CORSMiddleware sets CORS headers for every response.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow any origin to access the resource.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Allow these methods.
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// Allow these headers.
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		w.Header().
			Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With")
		// Expose specific headers to the client.
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		// If it's a preflight OPTIONS request, respond immediately.
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
