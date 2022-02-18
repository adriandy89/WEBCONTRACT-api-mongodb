package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/currencycontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// CurrencyRoutes => Rutas de monedas
func CurrencyRoutes(r *mux.Router) {
	r.HandleFunc("/api/currencies", middlewares.CheckDB(middlewares.ValidateJWT(currencycontroller.GetAllCurrencies))).Methods("GET")
	r.HandleFunc("/api/currency/new", middlewares.CheckDB(middlewares.ValidateJWT(currencycontroller.CurrencyRegister))).Methods("POST")
	r.HandleFunc("/api/currency/{id}", middlewares.CheckDB(middlewares.ValidateJWT(currencycontroller.UpdateCurrencyByID))).Methods("PUT")
	r.HandleFunc("/api/currency/{id}", middlewares.CheckDB(middlewares.ValidateJWT(currencycontroller.DeleteCurrencyByID))).Methods("DELETE")
}
