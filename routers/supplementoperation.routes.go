package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/supplementoperationcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// SupplementOperationRoutes
func SupplementOperationRoutes(r *mux.Router) {
	r.HandleFunc("/api/supplementoperations", middlewares.CheckDB(middlewares.ValidateJWT(supplementoperationcontroller.GetAllSupplementsOperations))).Methods("GET")
	r.HandleFunc("/api/supplementoperation/new", middlewares.CheckDB(middlewares.ValidateJWT(supplementoperationcontroller.SupplementOperationRegister))).Methods("POST")
	r.HandleFunc("/api/supplementoperation/{id}", middlewares.CheckDB(middlewares.ValidateJWT(supplementoperationcontroller.UpdateSupplementOperationByID))).Methods("PUT")
	r.HandleFunc("/api/supplementoperation/{id}", middlewares.CheckDB(middlewares.ValidateJWT(supplementoperationcontroller.DeletesupplementOperationByID))).Methods("DELETE")
}
