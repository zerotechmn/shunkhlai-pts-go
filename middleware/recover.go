package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(map[string]interface{}{
					"code":    500,
					"message": "Internal Server Error",
					"error":   "INTERNAL_ERROR",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
