package typefactcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"WEBCONTRACT-api-mongodb/services/typefactservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllTypeFacts
func GetAllTypeFacts(w http.ResponseWriter, r *http.Request) {

	eList, err := typefactservice.FindAllTypeFacts()
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

// TypeFactRegister
func TypeFactRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var typeFact models.TypeFact
		err := json.NewDecoder(r.Body).Decode(&typeFact)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var typeFactFounded bool = typefactservice.ValidateIfExistByDescriptionAndCode(typeFact.Description, typeFact.CodeFact)

		if typeFactFounded {
			errorservice.ErrorMessage(w, "Ya existe, verifique", 400)
			return
		}

		errr := typefactservice.InsertNewTypeFact(typeFact)
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

// UpdateTypeFactByID
func UpdateTypeFactByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var typeFact models.TypeFact
		err := json.NewDecoder(r.Body).Decode(&typeFact)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if typeFact.Description == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := typefactservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go typefactservice.CodeQuery(typeFact.CodeFact, c)
		go typefactservice.DescriptionQuery(typeFact.Description, n)

		if <-c && cUpdate.CodeFact != typeFact.CodeFact {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}
		if <-n && cUpdate.Description != typeFact.Description {
			errorservice.ErrorMessage(w, "Ya existe, Verifique", 400)
			return
		}

		count, err := typefactservice.UpdateByID(id, typeFact)
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

// DeleteTypeFactByID
func DeleteTypeFactByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := typefactservice.DeleteByID(id)
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
