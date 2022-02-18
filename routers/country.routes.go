package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/countrycontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// CountryRoutes => Rutas de provincias
func CountryRoutes(r *mux.Router) {
	r.HandleFunc("/api/countries", middlewares.CheckDB(middlewares.ValidateJWT(countrycontroller.GetAllCountries))).Methods("GET")
	r.HandleFunc("/api/country/new", middlewares.CheckDB(middlewares.ValidateJWT(countrycontroller.CountryRegister))).Methods("POST")
	r.HandleFunc("/api/country/{id}", middlewares.CheckDB(middlewares.ValidateJWT(countrycontroller.UpdateCountryByID))).Methods("PUT")
	r.HandleFunc("/api/country/{id}", middlewares.CheckDB(middlewares.ValidateJWT(countrycontroller.DeleteCountryByID))).Methods("DELETE")
}
