package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterDevelopersRoutes(r *mux.Router) {
	r.HandleFunc("developers/create", nil).Methods(http.MethodPost)
	r.HandleFunc("developers/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("developers/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("developers/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("developers", nil).Methods(http.MethodGet)
}
