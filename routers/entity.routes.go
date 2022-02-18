package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/entitycontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// EntityRoutes => Rutas de Entidades
func EntityRoutes(r *mux.Router) {
	r.HandleFunc("/api/entities", middlewares.CheckDB(middlewares.ValidateJWT(entitycontroller.GetAllEntities))).Methods("GET")
	r.HandleFunc("/api/entity/new", middlewares.CheckDB(middlewares.ValidateJWT(entitycontroller.EntityRegister))).Methods("POST")
	r.HandleFunc("/api/entity/{id}", middlewares.CheckDB(middlewares.ValidateJWT(entitycontroller.UpdateEntityByID))).Methods("PUT")
	r.HandleFunc("/api/entity/{id}", middlewares.CheckDB(middlewares.ValidateJWT(entitycontroller.DeleteEntityByID))).Methods("DELETE")
}
