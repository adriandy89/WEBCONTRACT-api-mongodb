package contractspecificcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	contractspecificservice "WEBCONTRACT-api-mongodb/services/contractspecific.service.go"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// GetContractsSpecific => Contratos por cantidad orden y tipo
func GetContractsSpecific(w http.ResponseWriter, r *http.Request) {

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
	cList, total, founded := contractspecificservice.FindByCountAndSort(codeCompany, number, order, typ, numberPage)
	if total == 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var cListResp models.ContractSpecificReponse = models.ContractSpecificReponse{Total: total, ContractSpecificList: cList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}

// ContractSpecificRegister => controlador de la ruta de registro de contratos
func ContractSpecificRegister(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		var contract models.ContractSpecific
		err := json.NewDecoder(r.Body).Decode(&contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var cFounded bool = contractspecificservice.ValidateIfExistByCodeCompanyCodeContractAndCodeSpecific(contract.CodeCompany, contract.CodeContract, contract.CodeSpecific)

		if cFounded {
			errorservice.ErrorMessage(w, "Ese Codigo de Contrato ya existe", 400)
			return
		}

		errr := contractspecificservice.InsertNewContractSpecific(contract)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos", 500)
			return
		} else {
			messageservice.SuccesMessage(w, "Contrato Especifico creado correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// DeleteContractSpecificByID => eliminar mediante un id en los parametros
func DeleteContractSpecificByID(w http.ResponseWriter, r *http.Request) {

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		vars := mux.Vars(r)
		var idContract string = vars["id"]
		if len(idContract) == 0 {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		founded := contractspecificservice.DeleteByID(idContract)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Contrato Especifico eliminado correctamente", 200)
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

		var contract models.ContractSpecific
		err := json.NewDecoder(r.Body).Decode(&contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if contract.CodeSpecific == "" || contract.CodeContract == "" {
			errorservice.ErrorMessage(w, "Error, verifique Codigos del Contrato", 400)
			return
		}

		c := make(chan bool)
		go contractspecificservice.CodeSpecificQuery(contract.CodeSpecific, c)

		cUpdate, ext := contractspecificservice.FindByID(idContract)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		if <-c && contract.CodeSpecific != cUpdate.CodeSpecific {
			errorservice.ErrorMessage(w, "Ese Codigo Especifico de Contrato ya existe", 400)
			return
		}

		count, err := contractspecificservice.UpdateByID(idContract, contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		messageservice.SuccesMessage(w, "Contrato Especifico actualizado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acción", 401)
		return
	}
}

// GetCodeContractSpecific
func GetCodeContractSpecific(w http.ResponseWriter, r *http.Request) {

	var contract models.FindCodeContractSpecific
	err := json.NewDecoder(r.Body).Decode(&contract)
	if err != nil {
		errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
		return
	}

	var cFounded bool = contractspecificservice.ValidateIfExistByCodeCompanyAndCodeContract(contract.CodeCompany, contract.CodeContract)

	if !cFounded {
		errorservice.ErrorMessage(w, "Ese Codigo de Contrato no existe para su organizacion", 400)
		return
	}

	cList, founded := contractspecificservice.GetNewCodeContractSpecific(contract.CodeCompany, contract.CodeContract)

	if len(cList) == 0 {

		code := 1

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(code)
		return
	}
	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var arrN []int
	for i := 0; i < len(cList); i++ {
		arrT := strings.Split(cList[i], ". ")
		if len(arrT) > 1 {
			numb, err1 := strconv.Atoi(arrT[len(arrT)-1])
			if err1 == nil {
				arrN = append(arrN, numb)
			}
		}
	}
	if arrN != nil {
		arrN = removeDuplicateElement(arrN)
		quick_sort(&arrN, 0, len(arrN)-1)

		code := arrN[len(arrN)-1] + 1

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(code)
		return
	} else {
		code := 1

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(code)
		return
	}
}

// ------------------ UTILES----------------
func quick_sort(arreglo *[]int, izquierda, derecha int) []int {
	// separar valores menores a izquierda de pivot
	// y valores mayores a la derecha del pivot
	if izquierda < derecha {
		arr := *arreglo
		limite, origen := derecha, izquierda
		pivot := arr[derecha]
		derecha--

		for izquierda <= derecha {
			// buscar en izquierda numero mayor que pivot
			for izquierda < len(arr) && arr[izquierda] < pivot {
				izquierda++
			}
			// buscar en derecha numero menor que pivot
			for derecha >= 0 && arr[derecha] > pivot {
				derecha--
			}

			if izquierda <= derecha {
				// intercambiar encontrados
				swap(arreglo, izquierda, derecha)
				// aumentar valores izquierda
				izquierda++
				// decrementar valores derecha
				derecha--
			}

		}
		// termina separacion izquierda | derecha
		swap(arreglo, izquierda, limite)
		quick_sort(arreglo, origen, derecha)
		quick_sort(arreglo, izquierda, limite)
	}

	return *arreglo
}
func swap(arreglo *[]int, izquierda, derecha int) {
	// intercambiar valor izquierda con derecha
	arr := *arreglo
	temporal := arr[izquierda]
	arr[izquierda] = arr[derecha]
	arr[derecha] = temporal
}
func removeDuplicateElement(addrs []int) []int {
	result := make([]int, 0, len(addrs))
	temp := map[int]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
