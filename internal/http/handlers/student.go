package student

import (
	"log/slog"
	"net/http"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a student")
		w.Write([]byte("Welcome to Go Server Students Api"))
	}
}
