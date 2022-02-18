package organismcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"WEBCONTRACT-api-mongodb/services/organismservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllOrganims => todos los organismos
func GetAllOrganims(w http.ResponseWriter, r *http.Request) {

	organismList, err := organismservice.FindAllOrganisms()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(organismList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(organismList)
}

// OrganismRegister => controlador de la ruta de registro de provincia
func OrganismRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var organism models.Organism
		err := json.NewDecoder(r.Body).Decode(&organism)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var organismFounded bool = organismservice.ValidateIfExistByNameAndCode(organism.Name, organism.CodeOrganism)

		if organismFounded {
			errorservice.ErrorMessage(w, "Ese Organismo ya existe", 400)
			return
		}

		errr := organismservice.InsertNewOrganism(organism)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Organismo creado correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateOrganismByID
func UpdateOrganismByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var organism models.Organism
		err := json.NewDecoder(r.Body).Decode(&organism)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if organism.Name == "" || organism.CodeOrganism == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := organismservice.FindByID(id)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go organismservice.CodeQuery(organism.CodeOrganism, c)
		go organismservice.NameQuery(organism.Name, n)

		if <-c && cUpdate.CodeOrganism != organism.CodeOrganism {
			errorservice.ErrorMessage(w, "Ese Organismo ya existe", 400)
			return
		}
		if <-n && cUpdate.Name != organism.Name {
			errorservice.ErrorMessage(w, "Ese Organismo ya existe", 400)
			return
		}

		count, err := organismservice.UpdateByID(id, organism)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Organismo actualizado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteOrganismByID
func DeleteOrganismByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var id string = vars["id"]
		if len(id) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := organismservice.DeleteByID(id)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Organismo eliminado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
