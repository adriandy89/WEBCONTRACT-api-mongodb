package clientprovidercontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/clientproviderservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetClientProviders => Clientes y Provedores por cantidad orden y tipo
func GetClientProviders(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
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
	cpList, total, founded := clientproviderservice.FindByCountAndSort(number, order, typ, numberPage)
	if total == 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var cpListResp models.ClientProviderReponse = models.ClientProviderReponse{Total: total, ClientProviderList: cpList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cpListResp)
}

// GetClientProvidersByWord => Clientes y Provedores por cantidad orden y tipo
func GetClientProvidersByWord(w http.ResponseWriter, r *http.Request) {

	var word models.Word

	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	vars := mux.Vars(r)
	var count string = vars["count"]
	var order string = vars["order"]
	var typ string = vars["typ"]
	var page string = vars["page"]
	number, err := strconv.Atoi(count)
	numberPage, err1 := strconv.Atoi(page)
	if number <= 0 || numberPage <= 0 || err != nil || err1 != nil || order == "" || typ == "" || word.Word == "" {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}
	cpList, founded := clientproviderservice.FindByNameOrCode(number, order, typ, numberPage, word.Word)

	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var cpListResp models.ClientProviderReponse = models.ClientProviderReponse{ClientProviderList: cpList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cpListResp)
}

// ClientProviderRegister => controlador de la ruta de registro de cliente y/o provedor
func ClientProviderRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		var clientProvider models.ClientProvider
		err := json.NewDecoder(r.Body).Decode(&clientProvider)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var cpFounded bool = clientproviderservice.ValidateIfExistByCustId(clientProvider.CustId)

		if cpFounded {
			errorservice.ErrorMessage(w, "Ese Cliento y/o Proveedor ya existe", 400)
			return
		}

		errr := clientproviderservice.InsertNewClientProvider(clientProvider)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos", 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Cliento y/o Proveedor creado correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteClientProviderByID => eliminar mediante un id en los parametros
func DeleteClientProviderByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {
		vars := mux.Vars(r)
		var idClientProvider string = vars["id"]
		if len(idClientProvider) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := clientproviderservice.DeleteByID(idClientProvider)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Cliente y/o Proveedor eliminado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// UpdateClientProviderByID => actualiza un cliente o proveedor mediante un id en los parametros
func UpdateClientProviderByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" {

		vars := mux.Vars(r)
		var idClientProvider string = vars["id"]
		if len(idClientProvider) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		var clientProvider models.ClientProvider
		err := json.NewDecoder(r.Body).Decode(&clientProvider)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if clientProvider.CustId == "" || clientProvider.Name == "" {
			errorservice.ErrorMessage(w, "Error en la validacion de datos, verifique", 400)
			return
		}

		c := make(chan bool)
		go clientproviderservice.CustIdQuery(clientProvider.CustId, c)

		cpUpdate, ext := clientproviderservice.FindByID(idClientProvider)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		if <-c && clientProvider.CustId != cpUpdate.CustId {
			errorservice.ErrorMessage(w, "Esa categoria ya existe", 400)
			return
		}

		count, err := clientproviderservice.UpdateByID(idClientProvider, clientProvider)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Categoria actualizada correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}
