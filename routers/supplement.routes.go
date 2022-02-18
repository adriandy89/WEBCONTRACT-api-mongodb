package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/supplementcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// SupplementRoutes => Rutas de Suplementos
func SupplementRoutes(r *mux.Router) {
	r.HandleFunc("/api/supplement/new", middlewares.CheckDB(middlewares.ValidateJWT(supplementcontroller.SupplementRegister))).Methods("POST")
	r.HandleFunc("/api/supplement/{id}", middlewares.CheckDB(middlewares.ValidateJWT(supplementcontroller.DeleteSupplementByID))).Methods("DELETE")
	r.HandleFunc("/api/supplement/{id}", middlewares.CheckDB(middlewares.ValidateJWT(supplementcontroller.UpdateSupplementByID))).Methods("PUT")
	r.HandleFunc("/api/supplement/{codeCompany}/{codeContract}/{codeReeup}", middlewares.CheckDB(middlewares.ValidateJWT(supplementcontroller.GetAllSuplementByCodeCompanyContractReeup))).Methods("GET")
}
