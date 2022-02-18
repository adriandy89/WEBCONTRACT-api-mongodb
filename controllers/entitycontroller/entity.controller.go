package entitycontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/entityservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllEntities => todas las Entidades
func GetAllEntities(w http.ResponseWriter, r *http.Request) {

	eList, err := entityservice.FindAllEntities()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(eList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eList)
}

// EntityRegister
func EntityRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var entity models.Entity
		err := json.NewDecoder(r.Body).Decode(&entity)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var entityFounded bool = entityservice.ValidateIfExistByNameAndCode(entity.CompanyName, entity.CodeCompany)

		if entityFounded {
			errorservice.ErrorMessage(w, "Esa Entidad ya existe", 400)
			return
		}

		errr := entityservice.InsertNewEntity(entity)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Entidad creada correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateEntityByID
func UpdateEntityByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var entity models.Entity
		err := json.NewDecoder(r.Body).Decode(&entity)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if entity.CompanyName == "" || entity.CodeCompany == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := entityservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go entityservice.CodeQuery(entity.CodeCompany, c)
		go entityservice.NameQuery(entity.CompanyName, n)

		if <-c && cUpdate.CodeCompany != entity.CodeCompany {
			errorservice.ErrorMessage(w, "Esa Entidad ya existe", 400)
			return
		}
		if <-n && cUpdate.CompanyName != entity.CompanyName {
			errorservice.ErrorMessage(w, "Esa Entidad ya existe", 400)
			return
		}

		count, err := entityservice.UpdateByID(id, entity)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Entidad actualizada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteEntityByID
func DeleteEntityByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := entityservice.DeleteByID(id)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Entidad eliminada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
