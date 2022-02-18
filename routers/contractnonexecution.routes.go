package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/contractnonexecutioncontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// ContractRoutes => Rutas de Contratos
func ContractNonExecutionRoutes(r *mux.Router) {
	r.HandleFunc("/api/contractnonexecution/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(contractnonexecutioncontroller.GetContracts))).Methods("GET")
	r.HandleFunc("/api/contractnonexecution/new", middlewares.CheckDB(middlewares.ValidateJWT(contractnonexecutioncontroller.ContractRegister))).Methods("POST")
	r.HandleFunc("/api/contractnonexecution/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractnonexecutioncontroller.UpdateContractByID))).Methods("PUT")
	r.HandleFunc("/api/contractnonexecution/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractnonexecutioncontroller.DeleteContractByID))).Methods("DELETE")
}
