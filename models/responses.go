package models

// LoginReponse => Cuerpo de respuesta la momento de hacer login
type LoginReponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

// CategoryReponse => Cuerpo de respuesta la momento de devolver categorias
type CategoryReponse struct {
	Total        int64       `json:"total"`
	CategoryList []*Category `json:"categoryList"`
}

// ClientProviderReponse => Cuerpo de respuesta la momento de devolver clientes y proveedores
type ClientProviderReponse struct {
	Total              int64             `json:"total"`
	ClientProviderList []*ClientProvider `json:"clientProviderList"`
}

// ContractReponse => Cuerpo de respuesta la momento de devolver contratos
type ContractReponse struct {
	Total        int64       `json:"total"`
	ContractList []*Contract `json:"contractList"`
	Ending       int         `json:"ending"`
	Ended        int         `json:"ended"`
}

// OfferRequestReponse => Cuerpo de respuesta la momento de devolver ofertas
type OfferRequestReponse struct {
	Total            int64           `json:"total"`
	OfferRequestList []*OfferRequest `json:"offerRequestList"`
}

// ContractNonExecutionReponse => Cuerpo de respuesta la momento de devolver contractnonexecution
type ContractNonExecutionReponse struct {
	Total                    int64                   `json:"total"`
	ContractNonExecutionList []*ContractNonExecution `json:"contractNonExecutionList"`
}

// ContractReponse => Cuerpo de respuesta la momento de devolver contratos
type ContractSpecificReponse struct {
	Total                int64               `json:"total"`
	ContractSpecificList []*ContractSpecific `json:"contractSpecificList"`
}

// TypeContractResponse => Cuerpo de respuesta la momento de devolver tipos de contratos
type TypeContractResponse struct {
	TypeContractList []*TypeContract `json:"typeContractList"`
}

// TypeContractResponse => Cuerpo de respuesta la momento de devolver objetos del contrato
type ObjectContractResponse struct {
	ObjectContractList []*ObjectContract `json:"objectContractList"`
}

type SectorResponse struct {
	SectortList []*Sector `json:"sectortList"`
}

type TotalContractDetailReponse struct {
	Active   int64 `json:"active"`
	Inactive int64 `json:"inactive"`
}

type StadisticsDetailReponse struct {
	Ended   int64 `json:"ended"`
	OutTerm int   `json:"outTerm"`
	Active  int64 `json:"active"`
	Total   int64 `json:"total"`
}

type StadisticsDetailTypeCoisReponse struct {
	Cup   int64 `json:"cup"`
	Mlc   int64 `json:"mlc"`
	Ambas int64 `json:"ambas"`
	Total int64 `json:"total"`
}
