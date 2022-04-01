package handlers

import (
	config "WEBCONTRACT-api-mongodb/config_loader"
	"WEBCONTRACT-api-mongodb/routers"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//const FSPATH = "./frontend/"

// Handlers => En la funcion que maneja las peticiones
func Handlers() {

	router := mux.NewRouter()
	routers.AuthRoutes(router)
	routers.UserRoutes(router)
	routers.CategoryRoutes(router)
	routers.ClientProviderRoutes(router)
	routers.ContractRoutes(router)
	routers.ContractNonExecutionRoutes(router)
	routers.ContractSpecificRoutes(router)
	routers.CountryRoutes(router)
	routers.CurrencyRoutes(router)
	routers.DescriptionStateRoutes(router)
	routers.EntityRoutes(router)
	routers.NonEjecutionRoutes(router)
	routers.ObjectContractRoutes(router)
	routers.OfferRequestRoutes(router)
	routers.OrganismRoutes(router)
	routers.SupplementRoutes(router)
	routers.SupplementSpecificRoutes(router)
	routers.SupplementOperationRoutes(router)
	routers.TypeContractRoutes(router)
	routers.TypeFactRoutes(router)
	routers.SectorRoutes(router)
	routers.FilesRoutes(router)
	routers.StadisticRoutes(router)

	//router.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend"))).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(&withCustom404PageHandler{http.Dir("./frontend")})).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = config.ServerPORT
	}

	log.Println("Servidor Online, Puerto:", config.ServerPORT)

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

type withCustom404PageHandler struct{ fs http.FileSystem }

func (w4h *withCustom404PageHandler) Open(name string) (http.File, error) {
	f, err := w4h.fs.Open(name)
	if os.IsNotExist(err) {
		return w4h.fs.Open("index.html")
	}
	return f, err
}
