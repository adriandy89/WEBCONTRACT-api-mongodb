package objectcontractcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"WEBCONTRACT-api-mongodb/services/objectcontractservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllObjectContracts => todas las Entidades
func GetAllObjectContracts(w http.ResponseWriter, r *http.Request) {

	eList, err := objectcontractservice.FindAllObjectContracts()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(eList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	var cListResp models.ObjectContractResponse = models.ObjectContractResponse{ObjectContractList: eList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)
}

// ObjectContractRegister
func ObjectContractRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var objectcontract models.ObjectContract
		err := json.NewDecoder(r.Body).Decode(&objectcontract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var objectcontractFounded bool = objectcontractservice.ValidateIfExistByNameAndCode(objectcontract.Name, objectcontract.CodeObjectContract)

		if objectcontractFounded {
			errorservice.ErrorMessage(w, "Ya existe, verifique", 400)
			return
		}

		errr := objectcontractservice.InsertNewobjectContract(objectcontract)
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

// UpdateObjectContractByID
func UpdateObjectContractByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var objectcontract models.ObjectContract
		err := json.NewDecoder(r.Body).Decode(&objectcontract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if objectcontract.Name == "" || objectcontract.CodeObjectContract == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := objectcontractservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go objectcontractservice.CodeQuery(objectcontract.CodeObjectContract, c)
		go objectcontractservice.NameQuery(objectcontract.Name, n)

		if <-c && cUpdate.CodeObjectContract != objectcontract.CodeObjectContract {
			errorservice.ErrorMessage(w, "Ya existe, verifique", 400)
			return
		}
		if <-n && cUpdate.Name != objectcontract.Name {
			errorservice.ErrorMessage(w, "Ya existe, verifique", 400)
			return
		}

		count, err := objectcontractservice.UpdateByID(id, objectcontract)
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

// DeleteObjectContractByID
func DeleteObjectContractByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := objectcontractservice.DeleteByID(id)
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
