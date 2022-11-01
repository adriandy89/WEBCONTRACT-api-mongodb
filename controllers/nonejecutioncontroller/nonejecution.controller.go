package nonejecutioncontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	nonejecutionservice "WEBCONTRACT-api-mongodb/services/nonejecutionservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllNonEjecution => todos los motivos de No Ejecucion
func GetAllNonEjecution(w http.ResponseWriter, r *http.Request) {

	eList, err := nonejecutionservice.FindAllNonEjecutions()
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

// NonEjecutionRegister
func NonEjecutionRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var nonEjecution models.NonEjecution
		err := json.NewDecoder(r.Body).Decode(&nonEjecution)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var nonEjecutionFounded bool = nonejecutionservice.ValidateIfExistByDescriptionAndCode(nonEjecution.Description, nonEjecution.CodeNonExecution)

		if nonEjecutionFounded {
			errorservice.ErrorMessage(w, "Ya existe, verifique", 400)
			return
		}

		errr := nonejecutionservice.InsertNewNonEjecution(nonEjecution)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Creado correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdatenonEjecutionByID
func UpdatenonEjecutionByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var nonEjecution models.NonEjecution
		err := json.NewDecoder(r.Body).Decode(&nonEjecution)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if nonEjecution.Description == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := nonejecutionservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go nonejecutionservice.CodeQuery(nonEjecution.CodeNonExecution, c)
		go nonejecutionservice.DescriptionQuery(nonEjecution.Description, n)

		if <-c && cUpdate.CodeNonExecution != nonEjecution.CodeNonExecution {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}
		if <-n && cUpdate.Description != nonEjecution.Description {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}

		count, err := nonejecutionservice.UpdateByID(id, nonEjecution)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Actualizado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeletenonEjecutionByID
func DeletenonEjecutionByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := nonejecutionservice.DeleteByID(id)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Eliminado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
