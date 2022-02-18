package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/contractspecificcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// ContractRoutes => Rutas de Contratos
func ContractSpecificRoutes(r *mux.Router) {
	r.HandleFunc("/api/contractspecific/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(contractspecificcontroller.GetContractsSpecific))).Methods("GET")
	r.HandleFunc("/api/contractspecific/codespecific", middlewares.CheckDB(middlewares.ValidateJWT(contractspecificcontroller.GetCodeContractSpecific))).Methods("POST")
	r.HandleFunc("/api/contractspecific/new", middlewares.CheckDB(middlewares.ValidateJWT(contractspecificcontroller.ContractSpecificRegister))).Methods("POST")
	r.HandleFunc("/api/contractspecific/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractspecificcontroller.UpdateContractByID))).Methods("PUT")
	r.HandleFunc("/api/contractspecific/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractspecificcontroller.DeleteContractSpecificByID))).Methods("DELETE")
}
