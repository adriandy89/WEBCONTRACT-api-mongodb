package contractnonexecutioncontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/contractnonexecutionservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetContracts => Contratos por cantidad orden y tipo
func GetContracts(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]
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
	cList, total, founded := contractnonexecutionservice.FindByCountAndSort(codeCompany, number, order, typ, numberPage)
	if total == 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var cListResp models.ContractNonExecutionReponse = models.ContractNonExecutionReponse{Total: total, ContractNonExecutionList: cList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}

// ContractRegister => controlador de la ruta de registro de contratos
func ContractRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		var contract models.ContractNonExecution
		err := json.NewDecoder(r.Body).Decode(&contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		errr := contractnonexecutionservice.InsertNewContract(contract)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos", 500)
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

// DeleteContractByID => eliminar mediante un id en los parametros
func DeleteContractByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		vars := mux.Vars(r)
		var idContract string = vars["id"]
		if len(idContract) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := contractnonexecutionservice.DeleteByID(idContract)
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

// UpdateContractByID => actualiza un contrato mediante un id en los parametros
func UpdateContractByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {

		vars := mux.Vars(r)
		var idContract string = vars["id"]
		if len(idContract) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var contract models.ContractNonExecution
		err := json.NewDecoder(r.Body).Decode(&contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if contract.CodeContract == "" {
			errorservice.ErrorMessage(w, "Error, verifique Codigo de Contrato", 400)
			return
		}

		_, ext := contractnonexecutionservice.FindByID(idContract)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		count, err := contractnonexecutionservice.UpdateByID(idContract, contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Contrato actualizado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
