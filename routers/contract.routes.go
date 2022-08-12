package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/contractcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// ContractRoutes => Rutas de Contratos
func ContractRoutes(r *mux.Router) {
	r.HandleFunc("/api/contract/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsByWord))).Methods("POST")
	r.HandleFunc("/api/contract-vigent/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsVigentByWord))).Methods("POST")
	r.HandleFunc("/api/contract/newcode/{codeCompany}/{year}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetCodeContract))).Methods("GET")
	r.HandleFunc("/api/contractdays/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetDaysByDefaultContract))).Methods("GET")
	r.HandleFunc("/api/contract/totals/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.TotalContractByCodeCompanyQueryClasif))).Methods("GET")
	r.HandleFunc("/api/contract/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContracts))).Methods("GET")
	r.HandleFunc("/api/contract-vigent/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsVigent))).Methods("GET")
	r.HandleFunc("/api/contracts/excel/{codeCompany}/{filter}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsStadisticAllEXCEL))).Methods("GET")
	r.HandleFunc("/api/contracts/alert-excel/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsEndingEXCEL))).Methods("GET")
	r.HandleFunc("/api/contracts/alert/{codeCompany}/{type}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsEnding))).Methods("GET")
	r.HandleFunc("/api/contracts/alert/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsEndingSpecificDate))).Methods("POST")
	r.HandleFunc("/api/contract/new/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.ContractRegister))).Methods("POST")
	r.HandleFunc("/api/contract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.UpdateContractByID))).Methods("PUT")
	r.HandleFunc("/api/contract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.DeleteContractByID))).Methods("DELETE")
	r.HandleFunc("/api/contract/{codeCompany}/{count}/{order}/{typ}/{page}/{filter}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsStadistic))).Methods("GET")
	r.HandleFunc("/api/contract/stadistics/{codeCompany}/{filter}", middlewares.CheckDB(middlewares.ValidateJWT(contractcontroller.GetContractsStadisticAll))).Methods("GET")
}
