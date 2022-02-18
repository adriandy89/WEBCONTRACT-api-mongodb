package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/typecontractcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// TypeContractRoutes
func TypeContractRoutes(r *mux.Router) {
	r.HandleFunc("/api/typecontracts", middlewares.CheckDB(middlewares.ValidateJWT(typecontractcontroller.GetAllTypeContracts))).Methods("GET")
	r.HandleFunc("/api/typecontract/new", middlewares.CheckDB(middlewares.ValidateJWT(typecontractcontroller.TypeContractRegister))).Methods("POST")
	r.HandleFunc("/api/typecontract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(typecontractcontroller.UpdateTypeContractByID))).Methods("PUT")
	r.HandleFunc("/api/typecontract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(typecontractcontroller.DeleteTypeContractByID))).Methods("DELETE")
}
