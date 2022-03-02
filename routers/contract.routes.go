package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/contractcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// ContractRoutes => Rutas de Contratos
func ContractRoutes(r *mux.Router) {
	r.HandleFunc("/api/contract/newcode/{codeCompany}/{year}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetCodeContract))).Methods("GET")
	r.HandleFunc("/api/contract/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContracts))).Methods("GET")
	r.HandleFunc("/api/contract/new/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.ContractRegister))).Methods("POST")
	r.HandleFunc("/api/contract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.UpdateContractByID))).Methods("PUT")
	r.HandleFunc("/api/contract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.DeleteContractByID))).Methods("DELETE")
}
