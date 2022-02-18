package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/descriptionstatecontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// DescriptionStateRoutes => Rutas de Descripciones de estado
func DescriptionStateRoutes(r *mux.Router) {
	r.HandleFunc("/api/descriptionstates", middlewares.CheckDB(middlewares.ValidateJWT(descriptionstatecontroller.GetAllDescriptionsStates))).Methods("GET")
	r.HandleFunc("/api/descriptionstate/new", middlewares.CheckDB(middlewares.ValidateJWT(descriptionstatecontroller.DescriptionStateRegister))).Methods("POST")
	r.HandleFunc("/api/descriptionstate/{id}", middlewares.CheckDB(middlewares.ValidateJWT(descriptionstatecontroller.UpdateByID))).Methods("PUT")
	r.HandleFunc("/api/descriptionstate/{id}", middlewares.CheckDB(middlewares.ValidateJWT(descriptionstatecontroller.DeleteByID))).Methods("DELETE")
}
