package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/supplementspecificcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// SupplementRoutes => Rutas de Suplementos
func SupplementSpecificRoutes(r *mux.Router) {
	r.HandleFunc("/api/supplementspecific/new", middlewares.CheckDB(middlewares.ValidateJWT(supplementspecificcontroller.SupplementRegister))).Methods("POST")
	r.HandleFunc("/api/supplementspecific/{id}", middlewares.CheckDB(middlewares.ValidateJWT(supplementspecificcontroller.DeleteSupplementByID))).Methods("DELETE")
	r.HandleFunc("/api/supplementspecific/{id}", middlewares.CheckDB(middlewares.ValidateJWT(supplementspecificcontroller.UpdateSupplementByID))).Methods("PUT")
	r.HandleFunc("/api/supplementspecific/{codeCompany}/{codeContract}/{codeReeup}/{codeSpecific}", middlewares.CheckDB(middlewares.ValidateJWT(supplementspecificcontroller.GetAllSuplementByCodeCompanyContractReeup))).Methods("GET")
}
