package descriptionstatecontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/descriptionstateservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllDescriptionsStates
func GetAllDescriptionsStates(w http.ResponseWriter, r *http.Request) {

	dEList, err := descriptionstateservice.FindAllDescriptionsStates()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(dEList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dEList)
}

// DescriptionStateRegister
func DescriptionStateRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var dE models.DescriptionState
		err := json.NewDecoder(r.Body).Decode(&dE)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var dEFounded bool = descriptionstateservice.ValidateIfExistByDescriptionAndCode(dE.Description, dE.CodeDescriptionState)

		if dEFounded {
			errorservice.ErrorMessage(w, "Esa Descripcion de estado ya existe", 400)
			return
		}

		errr := descriptionstateservice.InsertNewDescriptionState(dE)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Descripcion de estado creada correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateByID
func UpdateByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var dE models.DescriptionState
		err := json.NewDecoder(r.Body).Decode(&dE)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if dE.Description == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := descriptionstateservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go descriptionstateservice.CodeQuery(dE.CodeDescriptionState, c)
		go descriptionstateservice.DescriptionQuery(dE.Description, n)

		if <-c && cUpdate.CodeDescriptionState != dE.CodeDescriptionState {
			errorservice.ErrorMessage(w, "Esa Descripcion de estado ya existe", 400)
			return
		}
		if <-n && cUpdate.Description != dE.Description {
			errorservice.ErrorMessage(w, "Esa Descripcion de estado ya existe", 400)
			return
		}

		count, err := descriptionstateservice.UpdateByID(id, dE)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Descripcion de estado actualizada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteByID
func DeleteByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := descriptionstateservice.DeleteByID(id)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Descripcion de estado eliminada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
