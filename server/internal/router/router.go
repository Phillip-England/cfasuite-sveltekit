package router

import (
	"cfasuite/internal/database"
	"cfasuite/internal/handler"
	"cfasuite/internal/httpcontext"
	"cfasuite/internal/middleware"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Router struct {
	Mux *http.ServeMux
	DB  database.DB
	// Templates *template.Template
}

func NewRouter(db database.DB) (*Router, error) {
	// templates, err := templates.ParseTemplates()
	// if err != nil {
	// 	return nil, err

	// }
	return &Router{
		Mux: http.NewServeMux(),
		DB:  db,
		// Templates: templates,
	}, nil
}

func (r *Router) Add(path string, handler handler.HandlerFunc, middleware ...middleware.MiddlewareFunc) {
	route := &Route{
		Mux:     r.Mux,
		Path:    path,
		Handler: handler,
		// templates:  r.Templates,
		Middleware: middleware,
		DB:         r.DB,
	}
	route.Handle()
}

func (r *Router) Serve(port string, message string) {
	fmt.Println(message)
	http.ListenAndServe(":"+port, r.Mux)
}

type Route struct {
	Mux        *http.ServeMux
	Path       string
	Handler    handler.HandlerFunc
	Templates  *template.Template
	Middleware []middleware.MiddlewareFunc
	DB         database.DB
}

func (r *Route) Handle() {
	r.Mux.HandleFunc(r.Path, func(w http.ResponseWriter, req *http.Request) {
		Chain(w, req, r)
	})
}

type ChainFunc func(w http.ResponseWriter, r *http.Request, router *Route)

func Chain(w http.ResponseWriter, r *http.Request, route *Route) {
	customContext := &httpcontext.Context{
		Context:   context.Background(),
		Templates: route.Templates,
		StartTime: time.Now(),
		DB:        route.DB,
	}
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
		return
	}
	for _, mw := range route.Middleware {
		err := mw(customContext, w, r)
		if err != nil {
			return
		}
	}
	route.Handler(customContext, w, r)
	middleware.Log(customContext, w, r)
}
