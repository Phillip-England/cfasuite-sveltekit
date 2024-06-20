package api

import (
	"cfasuite/internal/httpcontext"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type GeneralResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	GeneralResponse
}

type LoginClaims struct {
	SessionToken string `json:"session_token"`
	jwt.StandardClaims
}

func Login(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	email := loginRequest.Email
	password := loginRequest.Password
	if email == "" || password == "" {
		response := LoginResponse{
			GeneralResponse: GeneralResponse{
				Message: "all fields required",
			},
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	if email != os.Getenv("ADMIN_EMAIL") || password != os.Getenv("ADMIN_PASSWORD") {
		response := LoginResponse{
			GeneralResponse: GeneralResponse{
				Message: "invalid credentials",
			},
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		return
	}
	expirationtime := time.Now().Add(24 * time.Hour)
	claims := &LoginClaims{
		SessionToken: os.Getenv("ADMIN_SESSION_TOKEN"),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationtime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("ADMIN_JWT_SECRET")))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    tokenString,
		Expires:  expirationtime,
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	response := LoginResponse{
		GeneralResponse: GeneralResponse{
			Message: "login successful",
		},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func Auth(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	tokenString := cookie.Value
	claims := &LoginClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ADMIN_JWT_SECRET")), nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if !token.Valid {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(claims)
}

type LogoutResponse struct {
	Message string `json:"message"`
}

func Logout(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now(),
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
	response := LogoutResponse{
		Message: "logout successful",
	}
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

type NewCFARequest struct {
	Name   string `json:"name"`
	Number string `json:"number"`
}

type NewCFAResponse struct {
	GeneralResponse
}

func NewCFA(httpContext *httpcontext.Context, w http.ResponseWriter, r *http.Request) {
	db := httpContext.DB
	fmt.Println(db)
	var req NewCFARequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Name == "" || req.Number == "" {
		response := NewCFAResponse{
			GeneralResponse: GeneralResponse{
				Message: "all fields required",
			},
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	response := NewCFAResponse{
		GeneralResponse: GeneralResponse{
			Message: "cfa created",
		},
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
