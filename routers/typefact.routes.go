package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/typefactcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// TypeFactRoutes
func TypeFactRoutes(r *mux.Router) {
	r.HandleFunc("/api/typefacts", middlewares.CheckDB(middlewares.ValidateJWT(typefactcontroller.GetAllTypeFacts))).Methods("GET")
	r.HandleFunc("/api/typefact/new", middlewares.CheckDB(middlewares.ValidateJWT(typefactcontroller.TypeFactRegister))).Methods("POST")
	r.HandleFunc("/api/typefact/{id}", middlewares.CheckDB(middlewares.ValidateJWT(typefactcontroller.UpdateTypeFactByID))).Methods("PUT")
	r.HandleFunc("/api/typefact/{id}", middlewares.CheckDB(middlewares.ValidateJWT(typefactcontroller.DeleteTypeFactByID))).Methods("DELETE")
}
