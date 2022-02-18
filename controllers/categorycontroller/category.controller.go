package categorycontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/categoryservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCategories => function GetCategories por cantidad orden y tipo
func GetCategories(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var count string = vars["count"]
		var order string = vars["order"]
		var typ string = vars["typ"]
		var page string = vars["page"]
		number, err := strconv.Atoi(count)
		numberPage, err1 := strconv.Atoi(page)
		if number <= 0 || numberPage <= 0 || err != nil || err1 != nil || order == "" || typ == "" {
			errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
			return
		}
		categoryList, total, founded := categoryservice.FindByCountAndSort(number, order, typ, numberPage)
		if total == 0 {
			errorservice.ErrorMessage(w, "No hay datos", 400)
			return
		}
		if !founded {
			errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
			return
		}

		var categoryListResp models.CategoryReponse = models.CategoryReponse{Total: total, CategoryList: categoryList}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(categoryListResp)

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// GetAllCategories => todas las categorias
func GetAllCategories(w http.ResponseWriter, r *http.Request) {

	categoryList, err := categoryservice.FindAllCategories()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(categoryList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	var categoryListResp models.CategoryReponse = models.CategoryReponse{CategoryList: categoryList}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categoryListResp)
}

// CategoryRegister => controlador de la ruta de registro de categoria
func CategoryRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var category models.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var categoryFounded bool = categoryservice.ValidateIfExistByNameAndCode(category.Name, category.CodeCategory)

		if categoryFounded {
			errorservice.ErrorMessage(w, "Esa categoria ya existe", 400)
			return
		}

		errr := categoryservice.InsertNewCategory(category)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Categoria creada correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateCategoryByID => actualiza una categoria mediante un id en los parametros
func UpdateCategoryByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var idCategory string = vars["id"]
		if len(idCategory) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var category models.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if category.Name == "" || category.CodeCategory == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		categoryUpdate, ext := categoryservice.FindByID(idCategory)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go categoryservice.CodeQuery(category.CodeCategory, c)
		go categoryservice.NameQuery(category.Name, n)

		if <-c && categoryUpdate.CodeCategory != category.CodeCategory {
			errorservice.ErrorMessage(w, "Esa categoria ya existe", 400)
			return
		}
		if <-n && categoryUpdate.Name != category.Name {
			errorservice.ErrorMessage(w, "Esa categoria ya existe", 400)
			return
		}

		count, err := categoryservice.UpdateByID(idCategory, category)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Categoria actualizada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteCategoryByID => eliminar una sola categoria mediante un id en los parametros
func DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var idCategory string = vars["id"]
		if len(idCategory) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := categoryservice.DeleteByID(idCategory)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Categoria eliminada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
