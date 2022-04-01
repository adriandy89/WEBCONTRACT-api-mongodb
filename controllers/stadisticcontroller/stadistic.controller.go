package stadisticcontroller

import (
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/errorservice"
	"WEBCONTRACT-api-mongodb/services/stadisticservice"
	"WEBCONTRACT-api-mongodb/services/supplementservice"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetStadisticsContractsActiveInactiveOutdateTotal(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var codeCompany string = vars["codeCompany"]

	cList, founded := stadisticservice.FindActivesByCodeCompanyAndDate(codeCompany)

	if !founded {
		errorservice.ErrorMessage(w, "Parametros Invalidos", 400)
		return
	}

	for i := 0; i < len(cList); i++ {

		cList[i].Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(cList[i].CodeCompany, cList[i].CodeContract, cList[i].CodeReeup)
		truncated := false
		for j := 0; j < len(cList[i].Supplements); j++ {
			if cList[i].Supplements[j].ExpireAt != nil {
				if cList[i].Supplements[j].ExpireAt.After(time.Now().Add(-24*time.Hour)) && cList[i].Supplements[j].State == "Activo" {
					truncated = true
				}
			}
		}

		if truncated {
			cList[i] = cList[len(cList)-1] // Copy last element to index i.
			cList[len(cList)-1] = nil      // Erase last element (write zero value).
			cList = cList[:len(cList)-1]   // Truncate slice.
			i--
		}
	}
	outTerm := len(cList)
	active, inactive := stadisticservice.TotalContractByCodeCompanyQueryClasif(codeCompany)

	cListResp := models.StadisticsDetailReponse{OutTerm: outTerm, Active: active - int64(outTerm), Ended: inactive, Total: active + inactive}
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
