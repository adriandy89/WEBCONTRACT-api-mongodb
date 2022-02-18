package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/objectcontractcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// ObjectContract => Rutas de Objetivos de los Contratos
func ObjectContractRoutes(r *mux.Router) {
	r.HandleFunc("/api/objectcontracts", middlewares.CheckDB(middlewares.ValidateJWT(objectcontractcontroller.GetAllObjectContracts))).Methods("GET")
	r.HandleFunc("/api/objectcontract/new", middlewares.CheckDB(middlewares.ValidateJWT(objectcontractcontroller.ObjectContractRegister))).Methods("POST")
	r.HandleFunc("/api/objectcontract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(objectcontractcontroller.UpdateObjectContractByID))).Methods("PUT")
	r.HandleFunc("/api/objectcontract/{id}", middlewares.CheckDB(middlewares.ValidateJWT(objectcontractcontroller.DeleteObjectContractByID))).Methods("DELETE")
}
