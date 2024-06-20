package main

import (
	"cfasuite/internal/api"
	"cfasuite/internal/database"
	"cfasuite/internal/handler"
	"cfasuite/internal/middleware"
	"cfasuite/internal/router"
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	_ = godotenv.Load()

	dbConnection, err := sqlx.Connect("sqlite3", "main.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer dbConnection.Close()

	db := database.SQLiteDB{DB: dbConnection}

	err = db.CreateTables()
	if err != nil {
		log.Fatalln(err)
	}

	err = db.PrintTables()
	if err != nil {
		log.Fatalln(err)
	}

	r, err := router.NewRouter(db)
	if err != nil {
		panic(err)
	}

	r.Add("GET /favicon.ico", handler.Favicon)
	r.Add("GET /static/", handler.Static)

	r.Add("POST /api/auth/login", api.Login, middleware.JSON)
	r.Add("POST /api/auth/logout", api.Logout, middleware.JSON)
	r.Add("GET /api/auth", api.Auth, middleware.JSON)

	r.Add("POST /api/cfa", api.NewCFA, middleware.JSON)

	port := "8080"
	corsMux := middleware.CorsHandler(r.Mux)
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(":"+port, corsMux)

}
