package middleware

import (
	"cfasuite/internal/httpcontext"
	"fmt"
	"net/http"
	"os"
	"time"
)

type MiddlewareFunc func(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error

func Log(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error {
	elapsedTime := time.Since(ctx.StartTime)
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("[%s][%s][%s][%s]\n", r.Method, r.URL.Path, formattedTime, elapsedTime)
	return nil
}

func IsNotGuest(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}
	if cookie.Value == os.Getenv("ADMIN_SESSION_TOKEN") {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}
	return nil
}

func CorsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", os.Getenv("CLIENT_URL"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func JSON(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	return nil
}
