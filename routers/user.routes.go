package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/usercontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// UserRoutes => Rutas de Usuarios
func UserRoutes(r *mux.Router) {
	r.HandleFunc("/api/user/new", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.UserRegister))).Methods("POST")
	r.HandleFunc("/api/user/{id}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.GetUserByID))).Methods("GET")
	r.HandleFunc("/api/user/{id}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.DeleteUserByID))).Methods("DELETE")
	r.HandleFunc("/api/user/{id}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.UpdateUserByID))).Methods("PUT")
	r.HandleFunc("/api/users/{company}", middlewares.CheckDB(middlewares.ValidateJWT(usercontroller.GetAllUsersByComapany))).Methods("GET")
}
