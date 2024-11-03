package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterFloorsRoutes(r *mux.Router) {
	r.HandleFunc("floors/create", nil).Methods(http.MethodPost)
	r.HandleFunc("floors/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("floors/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("floors/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("floors", nil).Methods(http.MethodGet)
}
