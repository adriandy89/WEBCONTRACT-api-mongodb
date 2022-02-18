package typecontractcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"WEBCONTRACT-api-mongodb/services/typecontractservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllTypeContracts
func GetAllTypeContracts(w http.ResponseWriter, r *http.Request) {

	eList, err := typecontractservice.FindAllTypeContracts()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(eList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	var cListResp models.TypeContractResponse = models.TypeContractResponse{TypeContractList: eList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)
}

// TypeContractRegister
func TypeContractRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var typeContract models.TypeContract
		err := json.NewDecoder(r.Body).Decode(&typeContract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var typeContractFounded bool = typecontractservice.ValidateIfExistByNameAndCode(typeContract.Name, typeContract.CodeTypeContract)

		if typeContractFounded {
			errorservice.ErrorMessage(w, "Ya existe, verifique", 400)
			return
		}

		errr := typecontractservice.InsertNewTypeContract(typeContract)
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

// UpdateTypeContractByID
func UpdateTypeContractByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var typeContract models.TypeContract
		err := json.NewDecoder(r.Body).Decode(&typeContract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if typeContract.Name == "" || typeContract.CodeTypeContract == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := typecontractservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go typecontractservice.CodeQuery(typeContract.CodeTypeContract, c)
		go typecontractservice.NameQuery(typeContract.Name, n)

		if <-c && cUpdate.CodeTypeContract != typeContract.CodeTypeContract {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}
		if <-n && cUpdate.Name != typeContract.Name {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}

		count, err := typecontractservice.UpdateByID(id, typeContract)
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

// DeleteTypeContractByID
func DeleteTypeContractByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := typecontractservice.DeleteByID(id)
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
