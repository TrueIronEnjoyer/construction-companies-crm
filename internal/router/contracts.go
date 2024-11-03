package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterContractsRoutes(r *mux.Router) {
	r.HandleFunc("contracts/create", nil).Methods(http.MethodPost)
	r.HandleFunc("contracts/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("contracts/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("contracts/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("contracts", nil).Methods(http.MethodGet)
}
