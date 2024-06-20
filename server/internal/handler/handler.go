package handler

import (
	"cfasuite/internal/httpcontext"
	"net/http"
	"path/filepath"
)

type HandlerFunc func(ctx *httpcontext.Context, w http.ResponseWriter, r *http.Request)

func Favicon(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	filePath := "favicon.ico"
	fullPath := filepath.Join(".", ".", filePath)
	http.ServeFile(w, r, fullPath)
}

func Static(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	filePath := r.URL.Path[len("/static/"):]
	fullPath := filepath.Join(".", "static", filePath)
	http.ServeFile(w, r, fullPath)
}

func Preflight(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
}
