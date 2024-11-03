package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", nil).Methods(http.MethodPost)
	r.HandleFunc("/signup", nil).Methods(http.MethodPost)

	r.HandleFunc("developers/create", nil).Methods(http.MethodPost)
	r.HandleFunc("developers/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("developers/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("developers/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("developers", nil).Methods(http.MethodGet)

	r.HandleFunc("employees/create", nil).Methods(http.MethodPost)
	r.HandleFunc("employees/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("employees/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("employees/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("employees", nil).Methods(http.MethodGet)

	r.HandleFunc("buyers/create", nil).Methods(http.MethodPost)
	r.HandleFunc("buyers/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("buyers/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("buyers/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("buyers", nil).Methods(http.MethodGet)

	r.HandleFunc("complexes/create", nil).Methods(http.MethodPost)
	r.HandleFunc("complexes/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("complexes/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("complexes/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("complexes", nil).Methods(http.MethodGet)

	r.HandleFunc("blocks/create", nil).Methods(http.MethodPost)
	r.HandleFunc("blocks/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("blocks/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("blocks/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("blocks", nil).Methods(http.MethodGet)

	r.HandleFunc("houses/create", nil).Methods(http.MethodPost)
	r.HandleFunc("houses/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("houses/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("houses/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("houses", nil).Methods(http.MethodGet)

	r.HandleFunc("floors/create", nil).Methods(http.MethodPost)
	r.HandleFunc("floors/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("floors/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("floors/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("floors", nil).Methods(http.MethodGet)

	r.HandleFunc("entrances/create", nil).Methods(http.MethodPost)
	r.HandleFunc("entrances/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("entrances/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("entrances/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("entrances", nil).Methods(http.MethodGet)

	r.HandleFunc("apartments/create", nil).Methods(http.MethodPost)
	r.HandleFunc("apartments/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("apartments/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("apartments/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("apartments", nil).Methods(http.MethodGet)

	r.HandleFunc("contracts/create", nil).Methods(http.MethodPost)
	r.HandleFunc("contracts/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("contracts/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("contracts/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("contracts", nil).Methods(http.MethodGet)
}
