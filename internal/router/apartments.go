package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterApartmentsRoutes(r *mux.Router) {
	r.HandleFunc("apartments/create", nil).Methods(http.MethodPost)
	r.HandleFunc("apartments/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("apartments/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("apartments/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("apartments", nil).Methods(http.MethodGet)
}
