package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/authcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// AuthRoutes => Rutas de validacion y autenticacion
func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/api/auth/login", middlewares.CheckDB(authcontroller.Login)).Methods("POST")
}
