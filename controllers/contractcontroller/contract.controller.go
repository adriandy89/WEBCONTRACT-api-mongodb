package contractcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/contractservice"
	contractspecificservice "WEBCONTRACT-api-mongodb/services/contractspecific.service.go"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	offerrequestservice "WEBCONTRACT-api-mongodb/services/offerRequestservice"
	"WEBCONTRACT-api-mongodb/services/supplementservice"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	cList, total, founded := contractservice.FindByCountAndSort(codeCompany, number, order, typ, numberPage)
	if total == 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}
	for i := 0; i < len(cList); i++ {
		cList[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cList[i].CodeCompany, cList[i].CodeContract, cList[i].CodeReeup)
		cList[i].Specifics, _ = contractspecificservice.FindAllByCodeCompanyContractReeup(cList[i].CodeCompany, cList[i].CodeContract, cList[i].CodeReeup)
		if cList[i].CodeOfert != "" {
			cList[i].Offer, _ = offerrequestservice.FindOneByCompanyReeupAndOffer(cList[i].CodeCompany, cList[i].CodeReeup, cList[i].CodeOfert)
		}

	}

	var cListResp models.ContractReponse = models.ContractReponse{Total: total, ContractList: cList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}

func GetContractsByWord(w http.ResponseWriter, r *http.Request) {

	var word models.Word

	err := json.NewDecoder(r.Body).Decode(&word)
	if err != nil {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]
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
	cList, total, founded := contractservice.FindByNameOrCode(codeCompany, number, order, typ, numberPage, word.Word)

	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var cListResp models.ContractReponse = models.ContractReponse{Total: total, ContractList: cList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)
}

// ContractRegister => controlador de la ruta de registro de contratos
func ContractRegister(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	var rol string = r.Header.Get("rol")
	if rol == "Admin" || rol == "SA" || rol == "Gestionador" {
		var contract models.Contract
		err := json.NewDecoder(r.Body).Decode(&contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}

		var cFounded bool = contractservice.ValidateIfExistByCodeContract(contract.CodeContract, codeCompany)

		if cFounded {
			errorservice.ErrorMessage(w, "Ese Codigo de Contrato ya existe", 400)
			return
		}

		errr := contractservice.InsertNewContract(contract)
		if errr != nil {
			errorservice.ErrorMessage(w, "Error en registro en la base de datos", 500)
			return
		} else {

			if contract.CodeOfert != "" && contract.CodeReeup != "" {
				offer, exist := offerrequestservice.FindOneByCompanyReeupAndOffer(contract.CodeCompany, contract.CodeReeup, contract.CodeOfert)
				if exist {
					offer.State = "Inactivo"
					offerrequestservice.UpdateByID(offer.ID.Hex(), *offer)
				}
			}

			messageservice.SuccesMessage(w, "Contrato creado correctamente", 200)
			return
		}

	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acci??n", 401)
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

		founded := contractservice.DeleteByID(idContract)
		if !founded {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}
		messageservice.SuccesMessage(w, "Contrato eliminado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acci??n", 401)
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

		var contract models.Contract
		err := json.NewDecoder(r.Body).Decode(&contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
			return
		}
		if contract.CodeContract == "" {
			errorservice.ErrorMessage(w, "Error, verifique Codigo de Contrato", 400)
			return
		}

		c := make(chan bool)
		go contractservice.CodeContractQuery(contract.CodeContract, c)

		cUpdate, ext := contractservice.FindByID(idContract)
		if !ext {
			errorservice.ErrorMessage(w, "El id enviado no es valido", 400)
			return
		}

		if <-c && contract.CodeContract != cUpdate.CodeContract {
			errorservice.ErrorMessage(w, "Ese Codigo de Contrato ya existe", 400)
			return
		}

		count, err := contractservice.UpdateByID(idContract, contract)
		if err != nil {
			errorservice.ErrorMessage(w, "Error al actualizar la base de datos", 500)
			return
		}
		if count == 0 {
			messageservice.SuccesMessage(w, "No se modificaron ninguno de los campos", 202)
			return
		}
		if contract.CodeOfert != "" && contract.CodeReeup != "" {
			offer, exist := offerrequestservice.FindOneByCompanyReeupAndOffer(contract.CodeCompany, contract.CodeReeup, contract.CodeOfert)
			if exist {
				offer.State = "Inactivo"
				offerrequestservice.UpdateByID(offer.ID.Hex(), *offer)
			}
		}
		messageservice.SuccesMessage(w, "Contrato actualizado correctamente", 200)
	} else {
		errorservice.ErrorMessage(w, "No tiene suficientes permisos para esta acci??n", 401)
		return
	}
}

// GetCodeContract
func GetCodeContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]
	var year string = vars["year"]
	if codeCompany == "" || year == "" {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}
	cList, founded := contractservice.GetNewCodeContract(codeCompany, year)

	if len(cList) == 0 {

		arrNew := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(arrNew)
		return
	}
	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var arrN []int
	for i := 0; i < len(cList); i++ {
		arrT := strings.Split(cList[i], "-")
		if len(arrT) > 1 {
			numb, err1 := strconv.Atoi(arrT[len(arrT)-2])
			if err1 == nil {
				arrN = append(arrN, numb)
			}
		}
	}
	if arrN != nil {
		arrN = removeDuplicateElement(arrN)
		quick_sort(&arrN, 0, len(arrN)-1)

		var j int = 0
		var result []int
		for i := 1; i <= arrN[len(arrN)-1]; i++ {
			if arrN[j] == i {
				j++
			} else {
				result = append(result, i)
			}
		}

		add := arrN[len(arrN)-1]
		for len(result) < 20 {
			add++
			result = append(result, add)
		}
		if arrN[len(arrN)-1] > result[len(result)-1] {
			result = append(result, arrN[len(arrN)-1]+1)

		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)
		return
	} else {
		arrNew := [15]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(arrNew)
		return
	}
}

func GetContractsEnding(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]
	var total string = vars["type"]

	cList, founded := contractservice.FindByCodeCompanyAndDate(codeCompany)

	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}
	today := time.Now().Add(-24 * time.Hour)
	ending := 0
	ended := 0
	for i := 0; i < len(cList); i++ {

		cList[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cList[i].CodeCompany, cList[i].CodeContract, cList[i].CodeReeup)
		truncated := false
		for j := 0; j < len(cList[i].Supplements); j++ {
			if cList[i].Supplements[j].ExpireAt != nil {
				if cList[i].Supplements[j].ExpireAt.After(time.Now().Add(360*time.Hour)) && cList[i].Supplements[j].State == "Activo" {
					truncated = true
				}
			}
		}
		cList[i].Specifics, _ = contractspecificservice.FindAllByCodeCompanyContractReeup(cList[i].CodeCompany, cList[i].CodeContract, cList[i].CodeReeup)
		if cList[i].CodeOfert != "" {
			cList[i].Offer, _ = offerrequestservice.FindOneByCompanyReeupAndOffer(cList[i].CodeCompany, cList[i].CodeReeup, cList[i].CodeOfert)
		}

		if truncated {
			cList[i] = cList[len(cList)-1] // Copy last element to index i.
			cList[len(cList)-1] = nil      // Erase last element (write zero value).
			cList = cList[:len(cList)-1]   // Truncate slice.
			i--
		} else {
			if cList[i].Supplements != nil && len(cList[i].Supplements) > 0 {
				bigger := cList[i].ExpireAt
				for j := 0; j < len(cList[i].Supplements); j++ {
					if cList[i].Supplements[j].ExpireAt != nil {
						if cList[i].Supplements[j].State == "Activo" {
							if bigger.Before(*cList[i].Supplements[j].ExpireAt) {
								bigger = cList[i].Supplements[j].ExpireAt
							}
						}
					}
				}
				if bigger.Before(today) {
					ended++
				} else {
					ending++
				}
			} else {
				if cList[i].ExpireAt.Before(today) {
					ended++
				} else {
					ending++
				}
			}
		}

	}
	var cListResp models.ContractReponse
	if total == "full" {
		cListResp = models.ContractReponse{ContractList: cList, Ending: ending, Ended: ended}
	} else {
		cListResp = models.ContractReponse{Ending: ending, Ended: ended}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}

func GetContractsEndingSpecificDate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	var at models.At
	err := json.NewDecoder(r.Body).Decode(&at)
	if err != nil {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	cList, total, founded := contractservice.FindByCodeCompanyAndSpecificDate(codeCompany, at.At)
	if total == 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}
	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}
	for i := 0; i < len(cList); i++ {
		cList[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cList[i].CodeCompany, cList[i].CodeContract, cList[i].CodeReeup)
		cList[i].Specifics, _ = contractspecificservice.FindAllByCodeCompanyContractReeup(cList[i].CodeCompany, cList[i].CodeContract, cList[i].CodeReeup)
		if cList[i].CodeOfert != "" {
			cList[i].Offer, _ = offerrequestservice.FindOneByCompanyReeupAndOffer(cList[i].CodeCompany, cList[i].CodeReeup, cList[i].CodeOfert)
		}

	}

	var cListResp models.ContractReponse = models.ContractReponse{Total: total, ContractList: cList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}

func TotalContractByCodeCompanyQueryClasif(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	active, inactive := contractservice.TotalContractByCodeCompanyQueryClasif(codeCompany)

	var Totals models.TotalContractDetailReponse = models.TotalContractDetailReponse{Active: active, Inactive: inactive}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Totals)
}

// GetCodeContract
func GetDaysByDefaultContract(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	days := contractservice.GetDaysByDefaultContract(codeCompany)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(days)
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
