package router

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	RegisterAuthRoutes(r)

	RegisterDevelopersRoutes(r)
	RegisterEmployeesRoutes(r)
	RegisterBuyersRoutes(r)
	RegisterContractsRoutes(r)

	RegisterComplexesRoutes(r)
	RegisterBlocksRoutes(r)
	RegisterHousesRoutes(r)
	RegisterFloorsRoutes(r)
	RegisterEntrancesRoutes(r)
	RegisterApartmentsRoutes(r)
}
