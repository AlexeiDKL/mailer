package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"log/slog"
)

func SetMiddleware(next http.Handler) http.Handler {
	functions := [](func(next http.Handler) http.Handler){
		logging,
		panicRecovery,
	}

	slog.Debug("Set middleware", len(functions))
	fmt.Println("Set middleware", len(functions))

	for _, fn := range functions {
		next = fn(next)
	}
	return next
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)

		slog.Info("", req.Method, req.RequestURI, time.Since(start))
	})
}

func panicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				slog.Error(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, req)
	})
}
