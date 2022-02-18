package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/offerrequestcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// OfferRequestRoutes => Rutas de Ofertas
func OfferRequestRoutes(r *mux.Router) {
	r.HandleFunc("/api/offerrequest/new", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.OfferRequestRegister))).Methods("POST")
	r.HandleFunc("/api/offerrequest/{id}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.DeleteOfferRequestByID))).Methods("DELETE")
	r.HandleFunc("/api/offerrequest/{id}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.UpdateOfferRequestByID))).Methods("PUT")
	r.HandleFunc("/api/offerrequest/{company}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.GetAllOfferRequestsByComapany))).Methods("GET")
}
