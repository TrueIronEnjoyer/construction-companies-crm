package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterBuyersRoutes(r *mux.Router) {
	r.HandleFunc("buyers/create", nil).Methods(http.MethodPost)
	r.HandleFunc("buyers/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("buyers/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("buyers/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("buyers", nil).Methods(http.MethodGet)
}
