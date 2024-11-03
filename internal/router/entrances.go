package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterEntrancesRoutes(r *mux.Router) {
	r.HandleFunc("entrances/create", nil).Methods(http.MethodPost)
	r.HandleFunc("entrances/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("entrances/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("entrances/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("entrances", nil).Methods(http.MethodGet)
}
