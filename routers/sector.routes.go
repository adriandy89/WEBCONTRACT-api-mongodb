package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/sectorcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// SectorRoutes => Rutas de Sectores de las entidades
func SectorRoutes(r *mux.Router) {
	r.HandleFunc("/api/sectors", middlewares.CheckDB(middlewares.ValidateJWT(sectorcontroller.GetAllSectors))).Methods("GET")
	r.HandleFunc("/api/sector/new", middlewares.CheckDB(middlewares.ValidateJWT(sectorcontroller.SectorRegister))).Methods("POST")
	r.HandleFunc("/api/sector/{id}", middlewares.CheckDB(middlewares.ValidateJWT(sectorcontroller.UpdateSectorByID))).Methods("PUT")
	r.HandleFunc("/api/sector/{id}", middlewares.CheckDB(middlewares.ValidateJWT(sectorcontroller.DeletesectorByID))).Methods("DELETE")
}
