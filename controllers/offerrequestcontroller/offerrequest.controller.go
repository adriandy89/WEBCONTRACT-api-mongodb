package offerrequestcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	offerrequestservice "WEBCONTRACT-api-mongodb/services/offerRequestservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// OfferRequestRegister
func OfferRequestRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		var offerRequest models.OfferRequest
		err := json.NewDecoder(r.Body).Decode(&offerRequest)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		errr := offerrequestservice.InsertNewOfferRequest(offerRequest)
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

// GetAllOfferRequests
func GetAllOfferRequestsByComapany(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		vars := mux.Vars(r)
		var company string = vars["company"]
		if len(company) == 0 {
			errorservice.ErrorMessage(w, "Codecompany no valido", 400)
			return
		}

		offerRequestList, founded := offerrequestservice.FindAllByCompany(company)
		if !founded {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(offerRequestList)

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteOfferRequestByID
func DeleteOfferRequestByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := offerrequestservice.DeleteByID(id)
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

// UpdateOfferRequestByID
func UpdateOfferRequestByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var offerRequest models.OfferRequest
		err := json.NewDecoder(r.Body).Decode(&offerRequest)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		_, ext := offerrequestservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		count, err := offerrequestservice.UpdateByID(id, offerRequest)
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
