package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/clientprovidercontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// ClientProvider => Rutas de clientes y proveedores
func ClientProviderRoutes(r *mux.Router) {
	r.HandleFunc("/api/clientprovider/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(clientprovidercontroller.GetClientProvidersByWord))).Methods("POST")
	r.HandleFunc("/api/clientprovider/{count}/{order}/{typ}/{page}/full", middlewares.CheckDB(middlewares.ValidateJWT(clientprovidercontroller.GetClientProvidersByWordFullData))).Methods("POST")
	r.HandleFunc("/api/clientprovider/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(clientprovidercontroller.GetClientProviders))).Methods("GET")
	r.HandleFunc("/api/clientprovider/new", middlewares.CheckDB(middlewares.ValidateJWT(clientprovidercontroller.ClientProviderRegister))).Methods("POST")
	r.HandleFunc("/api/clientprovider/{id}", middlewares.CheckDB(middlewares.ValidateJWT(clientprovidercontroller.UpdateClientProviderByID))).Methods("PUT")
	r.HandleFunc("/api/clientprovider/{id}", middlewares.CheckDB(middlewares.ValidateJWT(clientprovidercontroller.DeleteClientProviderByID))).Methods("DELETE")
}
