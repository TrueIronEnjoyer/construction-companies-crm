package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(r *mux.Router) {
	r.HandleFunc("/register", nil).Methods(http.MethodPost)
	r.HandleFunc("/signup", nil).Methods(http.MethodPost)
}
