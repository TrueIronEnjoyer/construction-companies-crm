package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterEmployeesRoutes(r *mux.Router) {
	r.HandleFunc("employees/create", nil).Methods(http.MethodPost)
	r.HandleFunc("employees/{id:[0-9]+}", nil).Methods(http.MethodGet)
	r.HandleFunc("employees/{id:[0-9]+}", nil).Methods(http.MethodPatch)
	r.HandleFunc("employees/{id:[0-9]+}", nil).Methods(http.MethodDelete)
	r.HandleFunc("employees", nil).Methods(http.MethodGet)
}
