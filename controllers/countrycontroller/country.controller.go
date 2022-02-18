package countrycontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/countryservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllCountries => todas las provincias
func GetAllCountries(w http.ResponseWriter, r *http.Request) {

	countryList, err := countryservice.FindAllCountries()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(countryList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(countryList)
}

// CountryRegister => controlador de la ruta de registro de provincia
func CountryRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var country models.Country
		err := json.NewDecoder(r.Body).Decode(&country)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var countryFounded bool = countryservice.ValidateIfExistByNameAndCode(country.Name, country.CodeCountry)

		if countryFounded {
			errorservice.ErrorMessage(w, "Esa provincia ya existe", 400)
			return
		}

		errr := countryservice.InsertNewCountry(country)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Provincia creada correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateCountryByID => actualiza una provincia mediante un id en los parametros
func UpdateCountryByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var idCountry string = vars["id"]
		if len(idCountry) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var country models.Country
		err := json.NewDecoder(r.Body).Decode(&country)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if country.Name == "" || country.CodeCountry == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := countryservice.FindByID(idCountry)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go countryservice.CodeQuery(country.CodeCountry, c)
		go countryservice.NameQuery(country.Name, n)

		if <-c && cUpdate.CodeCountry != country.CodeCountry {
			errorservice.ErrorMessage(w, "Esa provincia ya existe", 400)
			return
		}
		if <-n && cUpdate.Name != country.Name {
			errorservice.ErrorMessage(w, "Esa provincia ya existe", 400)
			return
		}

		count, err := countryservice.UpdateByID(idCountry, country)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Provincia actualizada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteCountryByID => eliminar una sola provincia mediante un id en los parametros
func DeleteCountryByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var idCountry string = vars["id"]
		if len(idCountry) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := countryservice.DeleteByID(idCountry)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Provincia eliminada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
