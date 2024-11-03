package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterComplexesRoutes(r *mux.Router) {
	r.HandleFunc("complexes/create", nil).Methods(http.MethodPost)
	r.HandleFunc("complexes/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("complexes/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("complexes/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("complexes", nil).Methods(http.MethodGet)
}
