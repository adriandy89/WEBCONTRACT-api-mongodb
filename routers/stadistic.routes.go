package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/stadisticcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// ContractRoutes => Rutas de Contratos
func StadisticRoutes(r *mux.Router) {
	r.HandleFunc("/api/stadistics/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(stadisticcontroller.GetStadisticsContractsActiveInactiveOutdateTotal))).Methods("GET")
	r.HandleFunc("/api/servertime", middlewares.CheckDB(middlewares.ValidateJWT(stadisticcontroller.GetServerTime))).Methods("GET")
	r.HandleFunc("/api/stadisticsCoins/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(stadisticcontroller.TotalTypeCoisByCodeCompany))).Methods("GET")
	r.HandleFunc("/api/stadisticsCoins/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(stadisticcontroller.TotalTypeCoisByCodeCompanyXDate))).Methods("POST")
	r.HandleFunc("/api/stadisticsClientsActives/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(stadisticcontroller.GetContractsClientProviderName))).Methods("GET")
	r.HandleFunc("/api/entitiesgroup/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(stadisticcontroller.GetStadisticsEntities))).Methods("GET")
}
