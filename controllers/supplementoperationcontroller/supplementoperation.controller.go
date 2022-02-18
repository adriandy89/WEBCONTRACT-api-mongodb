package supplementoperationcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"WEBCONTRACT-api-mongodb/services/supplementoperationservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllSupplementsOperations
func GetAllSupplementsOperations(w http.ResponseWriter, r *http.Request) {

	eList, err := supplementoperationservice.FindAllSupplementsOperations()
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

// SupplementOperationRegister
func SupplementOperationRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var supplementOperation models.SupplementOperation
		err := json.NewDecoder(r.Body).Decode(&supplementOperation)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var supplementOperationFounded bool = supplementoperationservice.ValidateIfExistByDescriptionAndCode(supplementOperation.Description, supplementOperation.CodeOperation)

		if supplementOperationFounded {
			errorservice.ErrorMessage(w, "Ya existe, verifique", 400)
			return
		}

		errr := supplementoperationservice.InsertNewSupplementOperation(supplementOperation)
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

// UpdateSupplementOperationByID
func UpdateSupplementOperationByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var supplementOperation models.SupplementOperation
		err := json.NewDecoder(r.Body).Decode(&supplementOperation)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if supplementOperation.Description == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := supplementoperationservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go supplementoperationservice.CodeQuery(supplementOperation.CodeOperation, c)
		go supplementoperationservice.DescriptionQuery(supplementOperation.Description, n)

		if <-c && cUpdate.CodeOperation != supplementOperation.CodeOperation {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}
		if <-n && cUpdate.Description != supplementOperation.Description {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}

		count, err := supplementoperationservice.UpdateByID(id, supplementOperation)
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

// DeletesupplementOperationByID
func DeletesupplementOperationByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := supplementoperationservice.DeleteByID(id)
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
