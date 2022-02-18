package currencycontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/currencyservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllCurrencies => todas las Monedas
func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {

	currencyList, err := currencyservice.FindAllCurrencies()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(currencyList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(currencyList)
}

// CurrencyRegister => controlador de la ruta de registro de moneda
func CurrencyRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var currency models.Currency
		err := json.NewDecoder(r.Body).Decode(&currency)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var currencyFounded bool = currencyservice.ValidateIfExistByNameAndCode(currency.Name, currency.Currency)

		if currencyFounded {
			errorservice.ErrorMessage(w, "Esa Moneda ya existe", 400)
			return
		}

		errr := currencyservice.InsertNewCurrency(currency)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos"+errr.Error(), 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Moneda creada correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateCurrencyByID => actualiza una moneda mediante un id en los parametros
func UpdateCurrencyByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var idCurrency string = vars["id"]
		if len(idCurrency) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var currency models.Currency
		err := json.NewDecoder(r.Body).Decode(&currency)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if currency.Name == "" || currency.Currency == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		cUpdate, ext := currencyservice.FindByID(idCurrency)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		c := make(chan bool)
		n := make(chan bool)
		go currencyservice.CodeQuery(currency.Currency, c)
		go currencyservice.NameQuery(currency.Name, n)

		if <-c && cUpdate.Currency != currency.Currency {
			errorservice.ErrorMessage(w, "Esa Moneda ya existe", 400)
			return
		}
		if <-n && cUpdate.Name != currency.Name {
			errorservice.ErrorMessage(w, "Esa Moneda ya existe", 400)
			return
		}

		count, err := currencyservice.UpdateByID(idCurrency, currency)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Moneda actualizada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteCurrencyByID => eliminar una sola moneda mediante un id en los parametros
func DeleteCurrencyByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var idCurrency string = vars["id"]
		if len(idCurrency) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := currencyservice.DeleteByID(idCurrency)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Moneda eliminada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
