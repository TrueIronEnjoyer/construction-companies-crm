package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterHousesRoutes(r *mux.Router) {
	r.HandleFunc("houses/create", nil).Methods(http.MethodPost)
	r.HandleFunc("houses/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("houses/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("houses/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("houses", nil).Methods(http.MethodGet)
}
