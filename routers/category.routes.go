package routers

import (
	"WEBCONTRACT-api-mongodb/controllers/categorycontroller"
	"WEBCONTRACT-api-mongodb/middlewares"

	"github.com/gorilla/mux"
)

// CategoryRoutes => Rutas de Categorias
func CategoryRoutes(r *mux.Router) {
	r.HandleFunc("/api/category/{count}/{order}/{typ}/{page}", middlewares.CheckDB(middlewares.ValidateJWT(categorycontroller.GetCategories))).Methods("GET")
	r.HandleFunc("/api/categories", middlewares.CheckDB(middlewares.ValidateJWT(categorycontroller.GetAllCategories))).Methods("GET")
	r.HandleFunc("/api/category/new", middlewares.CheckDB(middlewares.ValidateJWT(categorycontroller.CategoryRegister))).Methods("POST")
	r.HandleFunc("/api/category/{id}", middlewares.CheckDB(middlewares.ValidateJWT(categorycontroller.UpdateCategoryByID))).Methods("PUT")
	r.HandleFunc("/api/category/{id}", middlewares.CheckDB(middlewares.ValidateJWT(categorycontroller.DeleteCategoryByID))).Methods("DELETE")
}
