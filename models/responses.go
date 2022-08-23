package models

import (
	"time"
)

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
	EndedClient   int64 `json:"endedClient"`
	OutTermClient int   `json:"outTermClient"`
	ActiveClient  int64 `json:"activeClient"`
	TotalClient   int64 `json:"totalClient"`
	EndedProv     int64 `json:"endedProv"`
	OutTermProv   int   `json:"outTermProv"`
	ActiveProv    int64 `json:"activeProv"`
	TotalProv     int64 `json:"totalProv"`
}
type StadisticsDetailEntitiesReponse struct {
	Entidad            string  `json:"Entidad"`
	CodeCompany        string  `json:"CodeCompany"`
	CodeFather         string  `json:"CodeFather"`
	OutTermClient      int     `json:"Vencidos_Cientes"`
	ActiveClient       int64   `json:"Vigentes_Clientes"`
	EndedClient        int64   `json:"Terminados_Clientes"`
	TotalClient        int64   `json:"Total_Historico_Clientes"`
	TotalContractYearC int     `json:"Total_Anno_Clientes"`
	AmmountMNYearC     float32 `json:"MN_Anno_Clientes"`
	AmmountMLCYearC    float32 `json:"MLC_Anno_Clientes"`
	AmmountMNC         float32 `json:"MN_Total_Clientes"`
	AmmountMLCC        float32 `json:"MLC_Total_Clientes"`
	OutTermProv        int     `json:"Vencidos_Proveedores"`
	ActiveProv         int64   `json:"Vigentes_Proveedores"`
	EndedProv          int64   `json:"Terminados_Proveedores"`
	TotalProv          int64   `json:"Total_Historico_Proveedores"`
	TotalContractYearP int     `json:"Total_Anno_Proveedores"`
	AmmountMNYearP     float32 `json:"MN_Anno_Proveedores"`
	AmmountMLCYearP    float32 `json:"MLC_Anno_Proveedores"`
	AmmountMNP         float32 `json:"MN_Total_Proveedores"`
	AmmountMLCP        float32 `json:"MLC_Total_Proveedores"`
	OutTerm            int     `json:"Vencido"`
	Active             int64   `json:"Vigentes"`
	Ended              int64   `json:"Terminados"`
	Total              int64   `json:"Total_Historico"`
	AmmountMN          float32 `json:"MN_Total"`
	AmmountMLC         float32 `json:"MLC_Total"`
}

type StadisticsDetailTypeCoisReponse struct {
	Cup   int64 `json:"cup"`
	Mlc   int64 `json:"mlc"`
	Ambas int64 `json:"ambas"`
	Total int64 `json:"total"`
}

type StadisticsTotalsByMonthsReponse struct {
	YearActual  []int64 `json:"yearActual"`
	YearActualS []int64 `json:"yearActualS"`
	YearA       int     `json:"yearA"`
	YearBefore  []int64 `json:"yearBefore"`
	YearBeforeS []int64 `json:"yearBeforeS"`
	YearB       int     `json:"yearB"`
}

type Tree struct {
	CodeCompany string `json:"codeCompany"`
	CodeFather  string `json:"codeFather"`
}

type ContractEXCELResponse struct {
	ContractList []*ContractEXCEL `json:"contractList"`
}

type ContractEXCEL struct {
	ClientProviderName string             `bson:"clientProviderName,omitempty" json:"Cliente/Proveedor,omitempty"`
	CodeContract       string             `bson:"codeContract,omitempty" json:"Codigo_Contract,omitempty"`
	CodeReeup          string             `bson:"codeReeup,omitempty" json:"Reeup,omitempty"`
	CodeCompany        string             `bson:"codeCompany,omitempty" json:"Entidad,omitempty"`
	State              string             `bson:"state,omitempty" json:"Estado,omitempty"`
	CodeTypeCoin       string             `bson:"codeTypeCoin,omitempty" json:"Monedas,omitempty"`
	ClientSupplier     string             `bson:"clientSupplier,omitempty" json:"Tipo,omitempty"`
	CodeTypeContract   string             `bson:"codeTypeContract,omitempty" json:"Tipo_Contracto,omitempty"`
	CodeObject         string             `bson:"codeObject,omitempty" json:"Objetivo,omitempty"`
	CodeCategory       string             `bson:"codeCategory,omitempty" json:"Categoria,omitempty"`
	CreatedAt          *time.Time         `bson:"createdAt,omitempty" json:"Creado,omitempty"`
	ExpireAt           *time.Time         `bson:"expireAt,omitempty" json:"Expira,omitempty"`
	AmmountMN          float32            `bson:"ammountMN,omitempty" json:"MN,omitempty"`
	AmmountCUC         float32            `bson:"ammountCUC,omitempty" json:"CUC,omitempty"`
	AmmountMLC         float32            `bson:"ammountMLC,omitempty" json:"MLC,omitempty"`
	NonCompliance      []NonCompliance    `bson:"nonCompliance,omitempty" json:"Incumplimientos,omitempty"`
	Supplements        []*SupplementEXCEL `bson:"supplements,omitempty" json:"supplements,omitempty"`
}

type SupplementEXCEL struct {
	CodeContract   string     `bson:"codeContract,omitempty" json:"Codigo_Contract,omitempty"`
	CodeSupplement string     `bson:"codeSupplement,omitempty" json:"Codigo_Suplemento,omitempty"`
	CodeReeup      string     `bson:"codeReeup,omitempty" json:"Reeup,omitempty"`
	State          string     `bson:"state,omitempty" json:"Estado,omitempty"`
	OperationCUC   string     `bson:"operationCUC,omitempty" json:"Operacion,omitempty"`
	CreatedAt      *time.Time `bson:"createdAt,omitempty" json:"Creado,omitempty"`
	ExpireAt       *time.Time `bson:"expireAt,omitempty" json:"Expira,omitempty"`
	AmmountMN      float32    `bson:"ammountMN,omitempty" json:"MN,omitempty"`
	AmmountCUC     float32    `bson:"ammountCUC,omitempty" json:"CUC,omitempty"`
	AmmountMLC     float32    `bson:"ammountMLC,omitempty" json:"MLC,omitempty"`
	Description    string     `bson:"description,omitempty" json:"Descripcion,omitempty"`
}
