package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/filescontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// FilesRoutes => Rutas de subir y bajar Archivos
func FilesRoutes(r *mux.Router) {
	r.HandleFunc("/api/docs/{id}", middlewares.CheckDB(middlewares.ValidateJWT(filescontroller.DownloadFile))).Methods("GET")
	r.HandleFunc("/api/docs", middlewares.CheckDB(filescontroller.UploadFile)).Methods("POST")
	r.HandleFunc("/api/file/delete", middlewares.CheckDB(middlewares.ValidateJWT(filescontroller.DeleteFile))).Methods("POST")
}
