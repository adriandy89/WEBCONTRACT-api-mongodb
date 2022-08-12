package supplementcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"WEBCONTRACT-api-mongodb/services/supplementservice"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// SupplementRegister => controlador de la ruta de suplementos
func SupplementRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {

		var supplement models.Supplement
		err := json.NewDecoder(r.Body).Decode(&supplement)
		if err != nil {
			fmt.Println(r.Body)
			fmt.Println(err)
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		if supplementservice.ValidateIfExistByCodeContractAndCodeSupplement(supplement.CodeCompany, supplement.CodeContract, supplement.CodeReeup, supplement.CodeSupplement) {
			errorservice.ErrorMessage(w, "Codigo de Suplemento ya existe.", 400)
			return
		}

		errr := supplementservice.InsertNewSuplement(supplement)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Suplemento creado correctamente", 200)
			return
		}
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// GetAllSuplementByCodeCompanyContractReeup
func GetAllSuplementByCodeCompanyContractReeup(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]
	var codeContract string = vars["codeContract"]
	var codeReeup string = vars["codeReeup"]
	if len(codeContract) == 0 {
		errorservice.ErrorMessage(w, "codeContract no valido", 400)
		return
	}

	userList, founded := supplementservice.FindAllByCodeCompanyContractReeup(codeCompany, codeContract, codeReeup)
	if !founded {
		errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userList)

}

// DeleteSupplementByID
func DeleteSupplementByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := supplementservice.DeleteByID(id)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Suplemento eliminado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateSupplementByID
func UpdateSupplementByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var supplement models.Supplement
		err := json.NewDecoder(r.Body).Decode(&supplement)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if supplement.CodeContract == "" || supplement.CodeSupplement == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := supplementservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		if supplementservice.ValidateIfExistByCodeContractAndCodeSupplement(supplement.CodeCompany, supplement.CodeContract, supplement.CodeReeup, supplement.CodeSupplement) && cUpdate.CodeSupplement != supplement.CodeSupplement {
			errorservice.ErrorMessage(w, "Ese Codigo de Suplemento ya existe", 400)
			return
		}

		count, err := supplementservice.UpdateByID(id, supplement)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Suplemento actualizado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
