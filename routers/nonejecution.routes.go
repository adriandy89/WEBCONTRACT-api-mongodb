package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/nonejecutioncontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// NonEjecutionRoutes => Rutas de motivos de No ejecucion
func NonEjecutionRoutes(r *mux.Router) {
	r.HandleFunc("/api/nonejecutions", middlewares.CheckDB(middlewares.ValidateJWT(nonejecutioncontroller.GetAllNonEjecution))).Methods("GET")
	r.HandleFunc("/api/nonejecution/new", middlewares.CheckDB(middlewares.ValidateJWT(nonejecutioncontroller.NonEjecutionRegister))).Methods("POST")
	r.HandleFunc("/api/nonejecution/{id}", middlewares.CheckDB(middlewares.ValidateJWT(nonejecutioncontroller.UpdatenonEjecutionByID))).Methods("PUT")
	r.HandleFunc("/api/nonejecution/{id}", middlewares.CheckDB(middlewares.ValidateJWT(nonejecutioncontroller.DeletenonEjecutionByID))).Methods("DELETE")
}
