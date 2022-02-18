package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/organismcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// OrganismRoutes => Rutas de Organismos
func OrganismRoutes(r *mux.Router) {
	r.HandleFunc("/api/organisms", middlewares.CheckDB(middlewares.ValidateJWT(organismcontroller.GetAllOrganims))).Methods("GET")
	r.HandleFunc("/api/organism/new", middlewares.CheckDB(middlewares.ValidateJWT(organismcontroller.OrganismRegister))).Methods("POST")
	r.HandleFunc("/api/organism/{id}", middlewares.CheckDB(middlewares.ValidateJWT(organismcontroller.UpdateOrganismByID))).Methods("PUT")
	r.HandleFunc("/api/organism/{id}", middlewares.CheckDB(middlewares.ValidateJWT(organismcontroller.DeleteOrganismByID))).Methods("DELETE")
}
