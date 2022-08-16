package stadisticcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/entityservice"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/messageservice"
	"WEBCONTRACT-api-mongodb/services/stadisticservice"
	"WEBCONTRACT-api-mongodb/services/supplementservice"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// GetServerTime => Obtener Hora del Servidor
func GetServerTime(w http.ResponseWriter, r *http.Request) {
	messageservice.SuccesMessage(w, time.Now().Local().String(), 200)
}

func GetStadisticsContractsActiveInactiveOutdateTotal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	cListClient, cListProv, founded := stadisticservice.FindActivesByCodeCompanyAndDate(codeCompany)

	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}
	//clientes
	for i := 0; i < len(cListClient); i++ {

		cListClient[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cListClient[i].CodeCompany, cListClient[i].CodeContract, cListClient[i].CodeReeup)
		truncated := false
		for j := 0; j < len(cListClient[i].Supplements); j++ {
			if cListClient[i].Supplements[j].ExpireAt != nil {
				if cListClient[i].Supplements[j].ExpireAt.After(time.Now().Add(-24*time.Hour)) && cListClient[i].Supplements[j].State == "Activo" {
					truncated = true
				}
			}
		}

		if truncated {
			cListClient[i] = cListClient[len(cListClient)-1] // Copy last element to index i.
			cListClient[len(cListClient)-1] = nil            // Erase last element (write zero value).
			cListClient = cListClient[:len(cListClient)-1]   // Truncate slice.
			i--
		}
	}
	outTermClient := len(cListClient)
	//Proveedores
	for i := 0; i < len(cListProv); i++ {

		cListProv[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cListProv[i].CodeCompany, cListProv[i].CodeContract, cListProv[i].CodeReeup)
		truncated := false
		for j := 0; j < len(cListProv[i].Supplements); j++ {
			if cListProv[i].Supplements[j].ExpireAt != nil {
				if cListProv[i].Supplements[j].ExpireAt.After(time.Now().Add(-24*time.Hour)) && cListProv[i].Supplements[j].State == "Activo" {
					truncated = true
				}
			}
		}

		if truncated {
			cListProv[i] = cListProv[len(cListProv)-1] // Copy last element to index i.
			cListProv[len(cListProv)-1] = nil          // Erase last element (write zero value).
			cListProv = cListProv[:len(cListProv)-1]   // Truncate slice.
			i--
		}
	}
	outTermProv := len(cListProv)
	//	active, inactive := stadisticservice.TotalContractByCodeCompanyQueryClasif(codeCompany)
	activeClient, inactiveClient, activeProv, inactiveProv := stadisticservice.TotalContractByCodeCompanyQueryClasif(codeCompany)

	cListResp := models.StadisticsDetailReponse{
		OutTermClient: outTermClient, ActiveClient: activeClient - int64(outTermClient), EndedClient: inactiveClient, TotalClient: activeClient + inactiveClient,
		OutTermProv: outTermProv, ActiveProv: activeProv - int64(outTermProv), EndedProv: inactiveProv, TotalProv: activeProv + inactiveProv,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}

func TotalTypeCoisByCodeCompany(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	cup, mlc, ambas := stadisticservice.TotalTypeCoisByCodeCompany(codeCompany)

	cListResp := models.StadisticsDetailTypeCoisReponse{Cup: cup, Mlc: mlc, Ambas: ambas, Total: cup + mlc + ambas}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}
func TotalTypeCoisByCodeCompanyXDate(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	var rangeDate models.TimeRange
	err := json.NewDecoder(r.Body).Decode(&rangeDate)
	if err != nil {
		errorservice.ErrorMessage(w, "Error en la validacion de datos", 400)
		return
	}

	cup, mlc, ambas := stadisticservice.TotalTypeCoisByCodeCompanyXDate(codeCompany, rangeDate.Start, rangeDate.End)

	cListResp := models.StadisticsDetailTypeCoisReponse{Cup: cup, Mlc: mlc, Ambas: ambas, Total: cup + mlc + ambas}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)

}

func GetContractsClientProviderName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	cList, founded := stadisticservice.FindActivesByCodeCompanyGroupBy(codeCompany)

	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	var cListResp models.ContractReponse = models.ContractReponse{ContractList: cList}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cListResp)
}

// Por entidades
func GetStadisticsEntities(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	//var codeCompany string = "150"

	eList, err := entityservice.FindAllEntitiesCodeCompany()
	if err != nil {
		errorservice.ErrorMessage(w, "Invalidos: "+err.Error(), 400)
		return
	}
	if len(eList) <= 0 {
		errorservice.ErrorMessage(w, "No hay datos", 400)
		return
	}

	arrg := make([]models.Tree, 0)
	for i := 0; i < len(eList); i++ {
		arrg = append(arrg, models.Tree{CodeCompany: eList[i].CodeCompany, CodeFather: eList[i].CodeFather})
	}
	companies := make([]string, 0)
	companies = append(companies, codeCompany)
	for j := 0; j < len(companies); j++ {
		for k := 0; k < len(arrg); k++ {
			if companies[j] == arrg[k].CodeFather {
				companies = append(companies, arrg[k].CodeCompany)
			}
		}
	}

	respt := make([]models.StadisticsDetailEntitiesReponse, 0)
	for l := 0; l < len(companies); l++ {
		cListResp := getStadisticsContractsActiveInactiveOutdateTotal(companies[l])
		cListResp.CodeCompany = companies[l]
		cListResp.Entidad = entityservice.FindCompanyName(companies[l])		
		cListResp.CodeFather = entityservice.FindCompanyCodeFather(companies[l])
		respt = append(respt, cListResp)
	}
	//var cListResp models.EntitiesResponse = models.EntitiesResponse{EntitiestList: arrg}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(respt)
}
func getStadisticsContractsActiveInactiveOutdateTotal(codeCompany string) models.StadisticsDetailEntitiesReponse {

	cListClient, cListProv, _ := stadisticservice.FindActivesByCodeCompanyAndDate(codeCompany)

	//clientes
	for i := 0; i < len(cListClient); i++ {

		cListClient[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cListClient[i].CodeCompany, cListClient[i].CodeContract, cListClient[i].CodeReeup)
		truncated := false
		for j := 0; j < len(cListClient[i].Supplements); j++ {
			if cListClient[i].Supplements[j].ExpireAt != nil {
				if cListClient[i].Supplements[j].ExpireAt.After(time.Now().Add(-24*time.Hour)) && cListClient[i].Supplements[j].State == "Activo" {
					truncated = true
				}
			}
		}

		if truncated {
			cListClient[i] = cListClient[len(cListClient)-1] // Copy last element to index i.
			cListClient[len(cListClient)-1] = nil            // Erase last element (write zero value).
			cListClient = cListClient[:len(cListClient)-1]   // Truncate slice.
			i--
		}
	}
	outTermClient := len(cListClient)
	//Proveedores
	for i := 0; i < len(cListProv); i++ {

		cListProv[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cListProv[i].CodeCompany, cListProv[i].CodeContract, cListProv[i].CodeReeup)
		truncated := false
		for j := 0; j < len(cListProv[i].Supplements); j++ {
			if cListProv[i].Supplements[j].ExpireAt != nil {
				if cListProv[i].Supplements[j].ExpireAt.After(time.Now().Add(-24*time.Hour)) && cListProv[i].Supplements[j].State == "Activo" {
					truncated = true
				}
			}
		}

		if truncated {
			cListProv[i] = cListProv[len(cListProv)-1] // Copy last element to index i.
			cListProv[len(cListProv)-1] = nil          // Erase last element (write zero value).
			cListProv = cListProv[:len(cListProv)-1]   // Truncate slice.
			i--
		}
	}
	totalMNFC, totalMNYearFC, totalMLCFC, totalMLCYearFC, totalContractYearC := GetContractsWithSupplTotalCoins(codeCompany, "Cliente")
	totalMNFP, totalMNYearFP, totalMLCFP, totalMLCYearFP, totalContractYearP := GetContractsWithSupplTotalCoins(codeCompany, "Proveedor")

	outTermProv := len(cListProv)
	//	active, inactive := stadisticservice.TotalContractByCodeCompanyQueryClasif(codeCompany)
	activeClient, inactiveClient, activeProv, inactiveProv := stadisticservice.TotalContractByCodeCompanyQueryClasif(codeCompany)

	cListResp := models.StadisticsDetailEntitiesReponse{
		OutTermClient: outTermClient, ActiveClient: activeClient - int64(outTermClient), EndedClient: inactiveClient, TotalClient: activeClient + inactiveClient,
		OutTermProv: outTermProv, ActiveProv: activeProv - int64(outTermProv), EndedProv: inactiveProv, TotalProv: activeProv + inactiveProv,
		AmmountMNC: totalMNFC, AmmountMLCC: totalMLCFC, AmmountMNYearC: totalMNYearFC, AmmountMLCYearC: totalMLCYearFC, TotalContractYearC: totalContractYearC,
		AmmountMNP: totalMNFP, AmmountMLCP: totalMLCFP, AmmountMNYearP: totalMNYearFP, AmmountMLCYearP: totalMLCYearFP, TotalContractYearP: totalContractYearP,
		OutTerm: outTermClient + outTermProv, Active: (activeClient - int64(outTermClient) + activeProv - int64(outTermProv)),
		Ended: inactiveClient + inactiveProv, Total: activeClient + inactiveClient + activeProv + inactiveProv,
		AmmountMN: totalMNFC + totalMNFP, AmmountMLC: totalMLCFC + totalMLCFP,
	}
	return cListResp
}

// -----------------------------------------------------------------
// calcular el importe total de un contrato
func GetContractsWithSupplTotalCoins(codeCompany string, typeContract string) (float32, float32, float32, float32, int) {

	cList, founded := stadisticservice.ContractsWithSupplTotalCoins(codeCompany, typeContract)

	if !founded {
		return 0, 0, 0, 0, 0
	}

	var totalMNF float32 = 0
	var totalMNYearF float32 = 0
	var totalMLCF float32 = 0
	var totalMLCYearF float32 = 0
	var totalContractYear int = 0
	for i := 0; i < len(cList); i++ {
		var totalMN float32 = 0
		var totalMNYear float32 = 0
		var totalMLC float32 = 0
		var totalMLCYear float32 = 0
		if cList[i].AmmountMN > 0 {
			totalMN += cList[i].AmmountMN
		}
		if cList[i].AmmountCUC > 0 {
			totalMN += (cList[i].AmmountCUC * 24)
		}
		if cList[i].AmmountMLC > 0 {
			totalMLC += cList[i].AmmountMLC
		}
		if cList[i].CreatedAt.Year() == time.Now().Year() {
			totalMNYear = totalMN
			totalMLCYear = totalMLC
			totalContractYear++
		}
		for j := 0; j < len(cList[i].Supplements); j++ {
			switch cList[i].Supplements[j].OperationCUC {
			case "Nuevo valor pactado":
				totalMN = 0
				totalMLC = 0
				if cList[i].Supplements[j].AmmountMN > 0 {
					totalMN += cList[i].Supplements[j].AmmountMN
				}
				if cList[i].Supplements[j].AmmountCUC > 0 {
					totalMN += (cList[i].Supplements[j].AmmountCUC * 24)
				}
				if cList[i].Supplements[j].AmmountMLC > 0 {
					totalMLC += cList[i].Supplements[j].AmmountMLC
				}
			case "Aumento de Precios":
				if cList[i].Supplements[j].AmmountMN > 0 {
					totalMN += cList[i].Supplements[j].AmmountMN
				}
				if cList[i].Supplements[j].AmmountCUC > 0 {
					totalMN += (cList[i].Supplements[j].AmmountCUC * 24)
				}
				if cList[i].Supplements[j].AmmountMLC > 0 {
					totalMLC += cList[i].Supplements[j].AmmountMLC
				}
			case "Diminución de Precios":
				if cList[i].Supplements[j].AmmountMN > 0 {
					totalMN -= cList[i].Supplements[j].AmmountMN
				}
				if cList[i].Supplements[j].AmmountCUC > 0 {
					totalMN -= (cList[i].Supplements[j].AmmountCUC * 24)
				}
				if cList[i].Supplements[j].AmmountMLC > 0 {
					totalMLC -= cList[i].Supplements[j].AmmountMLC
				}
			default:
				if cList[i].Supplements[j].AmmountMN > 0 {
					totalMN += cList[i].Supplements[j].AmmountMN
				}
				if cList[i].Supplements[j].AmmountCUC > 0 {
					totalMN += (cList[i].Supplements[j].AmmountCUC * 24)
				}
				if cList[i].Supplements[j].AmmountMLC > 0 {
					totalMLC += cList[i].Supplements[j].AmmountMLC
				}
			}
			if cList[i].Supplements[j].CreatedAt.Year() == time.Now().Year() {
				switch cList[i].Supplements[j].OperationCUC {
				case "Nuevo valor pactado":
					totalMNYear = 0
					totalMLCYear = 0
					if cList[i].Supplements[j].AmmountMN > 0 {
						totalMNYear += cList[i].Supplements[j].AmmountMN
					}
					if cList[i].Supplements[j].AmmountCUC > 0 {
						totalMNYear += (cList[i].Supplements[j].AmmountCUC * 24)
					}
					if cList[i].Supplements[j].AmmountMLC > 0 {
						totalMLCYear += cList[i].Supplements[j].AmmountMLC
					}
				case "Aumento de Precios":
					if cList[i].Supplements[j].AmmountMN > 0 {
						totalMNYear += cList[i].Supplements[j].AmmountMN
					}
					if cList[i].Supplements[j].AmmountCUC > 0 {
						totalMNYear += (cList[i].Supplements[j].AmmountCUC * 24)
					}
					if cList[i].Supplements[j].AmmountMLC > 0 {
						totalMLCYear += cList[i].Supplements[j].AmmountMLC
					}
				case "Diminución de Precios":
					if cList[i].Supplements[j].AmmountMN > 0 {
						totalMNYear -= cList[i].Supplements[j].AmmountMN
					}
					if cList[i].Supplements[j].AmmountCUC > 0 {
						totalMNYear -= (cList[i].Supplements[j].AmmountCUC * 24)
					}
					if cList[i].Supplements[j].AmmountMLC > 0 {
						totalMLCYear -= cList[i].Supplements[j].AmmountMLC
					}
				default:
					if cList[i].Supplements[j].AmmountMN > 0 {
						totalMNYear += cList[i].Supplements[j].AmmountMN
					}
					if cList[i].Supplements[j].AmmountCUC > 0 {
						totalMNYear += (cList[i].Supplements[j].AmmountCUC * 24)
					}
					if cList[i].Supplements[j].AmmountMLC > 0 {
						totalMLCYear += cList[i].Supplements[j].AmmountMLC
					}
				}
			}
		}
		totalMNF += totalMN
		totalMNYearF += totalMNYear
		totalMLCF += totalMLC
		totalMLCYearF += totalMLCYear
	}
	//var cListResp models.ContractReponse = models.ContractReponse{AmmountMN: totalMNF, AmmountMLC: totalMLCF, AmmountMNYear: totalMNYearF, AmmountMLCYear: totalMLCYearF}
	return totalMNF, totalMNYearF, totalMLCF, totalMLCYearF, totalContractYear
}
