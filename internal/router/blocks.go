package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterBlocksRoutes(r *mux.Router) {
	r.HandleFunc("blocks/create", nil).Methods(http.MethodPost)
	r.HandleFunc("blocks/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("blocks/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("blocks/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("blocks", nil).Methods(http.MethodGet)
}
