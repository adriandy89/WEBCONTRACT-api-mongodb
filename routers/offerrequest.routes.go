package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/offerrequestcontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// OfferRequestRoutes => Rutas de Ofertas
func OfferRequestRoutes(r *mux.Router) {
	r.HandleFunc("/api/offerrequest/new/{codeCompany}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.OfferRequestRegister))).Methods("POST")
	r.HandleFunc("/api/offerrequest/{id}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.DeleteOfferRequestByID))).Methods("DELETE")
	r.HandleFunc("/api/offerrequest/{id}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.UpdateOfferRequestByID))).Methods("PUT")
	r.HandleFunc("/api/offerrequest/{codeCompany}/{codeReeup}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.GetAllOfferRequestsByCompanyAndReeup))).Methods("GET")
	r.HandleFunc("/api/offerrequest/{codeCompany}/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(offerrequestcontroller.GetOfferRequest))).Methods("GET")
}
