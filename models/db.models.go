package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User => Estructura usuario
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username    string             `bson:"username" json:"username,omitempty"`
	Password    string             `bson:"password,omitempty" json:"password,omitempty"`
	CreatedAt   time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	ExpireAt    time.Time          `bson:"expireAt,omitempty" json:"expireAt,omitempty"`
	State       int                `bson:"state" json:"state"`
	CodeCompany string             `bson:"codeCompany" json:"codeCompany,omitempty"`
	Name        string             `bson:"name" json:"name,omitempty"`
	Rol         string             `bson:"rol" json:"rol,omitempty"`
	Environment int                `bson:"environment" json:"environment"`
	LoginCount  int                `bson:"loginCount,omitempty" json:"loginCount,omitempty"`
}

// Country => Estructura de las Provincias
type Country struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeCountry string             `bson:"codeCountry" json:"codeCountry"`
	Name        string             `bson:"name" json:"name"`
	Initial     string             `bson:"initial" json:"initial,omitempty"`
}

// Organism => Estructura de los Organismos
type Organism struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeOrganism string             `bson:"codeOrganism" json:"codeOrganism"`
	Name         string             `bson:"name" json:"name"`
	Initial      string             `bson:"initial,omitempty" json:"initial,omitempty"`
}

// TypeContract => Estructura de los tipos de Contratos
type TypeContract struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeTypeContract string             `bson:"codeTypeContract" json:"codeTypeContract"`
	Name             string             `bson:"name" json:"name"`
}

// ObjectContract => Estructura de los objetivos de los Contratos
type ObjectContract struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeObjectContract string             `bson:"codeObjectContract" json:"codeObjectContract"`
	Name               string             `bson:"name" json:"name"`
}

// Currency => Estructura de las Monedas
type Currency struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Currency string             `bson:"currency" json:"currency"`
	Name     string             `bson:"name" json:"name"`
}

// Category => Estructura de las Categorias
type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeCategory string             `bson:"codeCategory" json:"codeCategory"`
	Name         string             `bson:"name" json:"name"`
}

// Entity  => Estructura de las entidades
type Entity struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeCompany string             `bson:"codeCompany" json:"codeCompany"`
	Address     string             `bson:"address,omitempty" json:"address,omitempty"`
	Prefix      string             `bson:"prefix,omitempty" json:"prefix,omitempty"`
	CompanyName string             `bson:"companyName" json:"companyName"`
	CodeFather  string             `bson:"codeFather,omitempty" json:"codeFather,omitempty"`
}

// UserRol => Estructura de Roles de Usuarios
type UserRolType struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Type        string             `bson:"type" json:"type"`
	Description string             `bson:"description" json:"description,omitempty"`
}

// DescriptionState => Estructura de las opciones de finalizar contrato
type DescriptionState struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeDescriptionState int                `bson:"codeDescriptionState" json:"codeDescriptionState"`
	Description          string             `bson:"description" json:"description"`
}

// NonEjecution => Estructura de las opciones de Incumplimiento - Tipo de reclamaciones
type NonEjecution struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeNonExecution int                `bson:"codeNonExecution" json:"codeNonExecution"`
	Description      string             `bson:"description" json:"description"`
}

// OfferRequest => Estructura de las Solicitudes de Ofertas
type OfferRequest struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeReeup   string             `bson:"codeReeup" json:"codeReeup"`
	CodeOffer   string             `bson:"codeOffer" json:"codeOffer"`
	CreatedAt   *time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	AmmountMN   float32            `bson:"ammountMN" json:"ammountMN"`
	AmmountCUC  float32            `bson:"ammountCUC" json:"ammountCUC"`
	State       string             `bson:"state" json:"state"`
	Description string             `bson:"description" json:"description"`
	CodeCompany string             `bson:"codeCompany" json:"codeCompany"`
	FileRoute   []string           `bson:"fileRoute" json:"fileRoute"`
}

// SupplementOperation => Estructura de las diferentes operaciones de los suplementos
type SupplementOperation struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeOperation int                `bson:"codeOperation" json:"codeOperation"`
	Description   string             `bson:"description" json:"description"`
}

// Sector => Estructura de los sectores
type Sector struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeSector int                `bson:"codeSector" json:"codeSector"`
	Name       string             `bson:"name" json:"name"`
}

// Supplement => Estructura de los Suplementos al Contrato
type Supplement struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CodeContract   string             `bson:"codeContract,omitempty" json:"codeContract,omitempty"`
	CodeReeup      string             `bson:"codeReeup,omitempty" json:"codeReeup,omitempty"`
	CodeSupplement string             `bson:"codeSupplement,omitempty" json:"codeSupplement,omitempty"`
	CreatedAt      *time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	ExpireAt       *time.Time         `bson:"expireAt,omitempty" json:"expireAt,omitempty"`
	State          string             `bson:"state,omitempty" json:"state,omitempty"`
	Description    string             `bson:"description,omitempty" json:"description,omitempty"`
	AmmountMN      float32            `bson:"ammountMN,omitempty" json:"ammountMN,omitempty"`
	AmmountCUC     float32            `bson:"ammountCUC,omitempty" json:"ammountCUC,omitempty"`
	OperationMN    int                `bson:"operationMN,omitempty" json:"operationMN,omitempty"`
	OperationCUC   string             `bson:"operationCUC,omitempty" json:"operationCUC,omitempty"`
	CodeCompany    string             `bson:"codeCompany,omitempty" json:"codeCompany,omitempty"`
	FileRoute      []string           `bson:"fileRoute,omitempty" json:"fileRoute,omitempty"`
}

// SupplementSpecific => Estructura de los Suplementos al Contrato
type SupplementSpecific struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeSpecific   string             `bson:"codeSpecific" json:"codeSpecific"`
	CodeContract   string             `bson:"codeContract" json:"codeContract"`
	CodeReeup      string             `bson:"codeReeup" json:"codeReeup"`
	CodeSupplement string             `bson:"codeSupplement" json:"codeSupplement"`
	CreatedAt      *time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	ExpireAt       *time.Time         `bson:"expireAt,omitempty" json:"expireAt,omitempty"`
	State          string             `bson:"state" json:"state,omitempty"`
	Description    string             `bson:"description,omitempty" json:"description,omitempty"`
	AmmountMN      float32            `bson:"ammountMN,omitempty" json:"ammountMN,omitempty"`
	AmmountCUC     float32            `bson:"ammountCUC,omitempty" json:"ammountCUC,omitempty"`
	OperationMN    int                `bson:"operationMN,omitempty" json:"operationMN,omitempty"`
	OperationCUC   string             `bson:"operationCUC,omitempty" json:"operationCUC,omitempty"`
	CodeCompany    string             `bson:"codeCompany" json:"codeCompany"`
	FileRoute      []string           `bson:"fileRoute,omitempty" json:"fileRoute,omitempty"`
}

// TypeFact => Estructura de Tipo de facturacion
type TypeFact struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeFact    int                `bson:"codeFact" json:"codeFact"`
	Description string             `bson:"description" json:"description"`
}

// Contract => Estructura de los Contratos
type Contract struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CodeContract            string             `bson:"codeContract,omitempty" json:"codeContract,omitempty"`
	CodeReeup               string             `bson:"codeReeup,omitempty" json:"codeReeup,omitempty"`
	CodeCompany             string             `bson:"codeCompany,omitempty" json:"codeCompany,omitempty"`
	CreatedAt               *time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	ExpireAt                *time.Time         `bson:"expireAt,omitempty" json:"expireAt,omitempty"`
	CodeTypeCoin            string             `bson:"codeTypeCoin,omitempty" json:"codeTypeCoin,omitempty"`
	CodeTypeContract        string             `bson:"codeTypeContract,omitempty" json:"codeTypeContract,omitempty"`
	State                   string             `bson:"state,omitempty" json:"state,omitempty"`
	AmmountMN               float32            `bson:"ammountMN,omitempty" json:"ammountMN,omitempty"`
	AmmountCUC              float32            `bson:"ammountCUC,omitempty" json:"ammountCUC,omitempty"`
	VerdictLegal            string             `bson:"verdictLegal,omitempty" json:"verdictLegal,omitempty"`
	ProcessPersonI          string             `bson:"processPersonI,omitempty" json:"processPersonI,omitempty"`
	ProcessPersonF          string             `bson:"processPersonF,omitempty" json:"processPersonF,omitempty"`
	NroArchive              int                `bson:"nroArchive,omitempty" json:"nroArchive,omitempty"`
	PayPeriod               string             `bson:"payPeriod,omitempty" json:"payPeriod,omitempty"`
	CreditDays              int                `bson:"creditDays,omitempty" json:"creditDays,omitempty"`
	ClientSupplier          string             `bson:"clientSupplier,omitempty" json:"clientSupplier,omitempty"`
	DateSuplementary        *time.Time         `bson:"dateSuplementary,omitempty" json:"dateSuplementary,omitempty"`
	CodeOfert               string             `bson:"codeOfert,omitempty" json:"codeOfert,omitempty"`
	CodeObject              string             `bson:"codeObject,omitempty" json:"codeObject,omitempty"`
	AmmountMNSuplementary   float32            `bson:"ammountMNSuplementary,omitempty" json:"ammountMNSuplementary,omitempty"`
	AmmountCUCSuplementary  float32            `bson:"ammountCUCSuplementary,omitempty" json:"ammountCUCSuplementary,omitempty"`
	CodeDescriptionState    string             `bson:"codeDescriptionState,omitempty" json:"codeDescriptionState,omitempty"`
	CommentDescriptionState string             `bson:"commentDescriptionState,omitempty" json:"commentDescriptionState,omitempty"`
	NonCompliance           []NonCompliance    `bson:"nonCompliance,omitempty" json:"nonCompliance,omitempty"`
	PaymentTerm             int                `bson:"paymentTerm,omitempty" json:"paymentTerm,omitempty"`
	CodeCategory            string             `bson:"codeCategory,omitempty" json:"codeCategory,omitempty"`
	AmmountMNInit           float32            `bson:"ammountMNInit,omitempty" json:"ammountMNInit,omitempty"`
	FileRoute               []string           `bson:"fileRoute,omitempty" json:"fileRoute,omitempty"`
	ClientProviderName      string             `bson:"clientProviderName,omitempty" json:"clientProviderName,omitempty"`
	Supplements             []*Supplement      `bson:"supplements,omitempty" json:"supplements,omitempty"`
}
type NonCompliance struct {
	Name        string `bson:"name,omitempty" json:"name,omitempty"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

// ContractSpecific => Estructura de los Contratos Especificos
type ContractSpecific struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeSpecific            string             `bson:"codeSpecific" json:"codeSpecific"`
	CodeContract            string             `bson:"codeContract" json:"codeContract"`
	CodeReeup               string             `bson:"codeReeup,omitempty" json:"codeReeup,omitempty"`
	CodeCompany             string             `bson:"codeCompany,omitempty" json:"codeCompany,omitempty"`
	CreatedAt               *time.Time         `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	ExpireAt                *time.Time         `bson:"expireAt,omitempty" json:"expireAt,omitempty"`
	CodeTypeCoin            string             `bson:"codeTypeCoin,omitempty" json:"codeTypeCoin,omitempty"`
	CodeTypeContract        string             `bson:"codeTypeContract,omitempty" json:"codeTypeContract,omitempty"`
	State                   string             `bson:"state,omitempty" json:"state,omitempty"`
	AmmountMN               float32            `bson:"ammountMN,omitempty" json:"ammountMN,omitempty"`
	AmmountCUC              float32            `bson:"ammountCUC,omitempty" json:"ammountCUC,omitempty"`
	VerdictLegal            string             `bson:"verdictLegal,omitempty" json:"verdictLegal,omitempty"`
	ProcessPersonI          string             `bson:"processPersonI,omitempty" json:"processPersonI,omitempty"`
	ProcessPersonF          string             `bson:"processPersonF,omitempty" json:"processPersonF,omitempty"`
	NroArchive              int                `bson:"nroArchive,omitempty" json:"nroArchive,omitempty"`
	PayPeriod               string             `bson:"payPeriod,omitempty" json:"payPeriod,omitempty"`
	CreditDays              int                `bson:"creditDays,omitempty" json:"creditDays,omitempty"`
	ClientSupplier          string             `bson:"clientSupplier,omitempty" json:"clientSupplier,omitempty"`
	DateSuplementary        *time.Time         `bson:"dateSuplementary,omitempty" json:"dateSuplementary,omitempty"`
	CodeOfert               string             `bson:"codeOfert,omitempty" json:"codeOfert,omitempty"`
	CodeObject              string             `bson:"codeObject,omitempty" json:"codeObject,omitempty"`
	AmmountMNSuplementary   float32            `bson:"ammountMNSuplementary,omitempty" json:"ammountMNSuplementary,omitempty"`
	AmmountCUCSuplementary  float32            `bson:"ammountCUCSuplementary,omitempty" json:"ammountCUCSuplementary,omitempty"`
	CodeDescriptionState    string             `bson:"codeDescriptionState,omitempty" json:"codeDescriptionState,omitempty"`
	CommentDescriptionState string             `bson:"commentDescriptionState,omitempty" json:"commentDescriptionState,omitempty"`
	PaymentTerm             int                `bson:"paymentTerm,omitempty" json:"paymentTerm,omitempty"`
	CodeCategory            string             `bson:"codeCategory,omitempty" json:"codeCategory,omitempty"`
	AmmountMNInit           float32            `bson:"ammountMNInit,omitempty" json:"ammountMNInit,omitempty"`
	FileRoute               []string           `bson:"fileRoute,omitempty" json:"fileRoute,omitempty"`
}

// ContractNonExecution => Estructura de Contratos no ejecutados
type ContractNonExecution struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CodeContract           string             `bson:"codeContract" json:"codeContract"`
	CodeReeup              string             `bson:"codeReeup" json:"codeReeup"`
	IdNonExecution         string             `bson:"idNonExecution" json:"idNonExecution"`
	ContractNonExecution   string             `bson:"contractNonExecution" json:"contractNonExecution"`
	IdContractNonExecution int                `bson:"idContractNonExecution" json:"idContractNonExecution"`
	CodeCompany            string             `bson:"codeCompany" json:"codeCompany"`
}

// ContractNonExecution => Estructura de Contratos Especificos no ejecutados
type ContractNonExecutionSpecific struct {
	CodeSpecific           string `bson:"codeSpecific" json:"codeSpecific"`
	CodeContract           string `bson:"codeContract" json:"codeContract"`
	CodeReeup              string `bson:"codeReeup" json:"codeReeup"`
	IdNonExecution         string `bson:"idNonExecution" json:"idNonExecution"`
	ContractNonExecution   string `bson:"contractNonExecution" json:"contractNonExecution"`
	IdContractNonExecution int    `bson:"idContractNonExecution" json:"idContractNonExecution"`
	CodeCompany            string `bson:"codeCompany" json:"codeCompany"`
}

// ClientProvider => Estructura de Cliente o proveedor
type ClientProvider struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CustId     string             `bson:"custId,omitempty" json:"custId,omitempty"`
	CodOne     string             `bson:"codOne,omitempty" json:"codOne,omitempty"`
	Name       string             `bson:"name,omitempty" json:"name,omitempty"`
	Organism   string             `bson:"organism,omitempty" json:"organism,omitempty"`
	StatusCode string             `bson:"statusCode,omitempty" json:"statusCode,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	Phone      string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Country    string             `bson:"country,omitempty" json:"country,omitempty"`
	SectorType string             `bson:"sectorType,omitempty" json:"sectorType,omitempty"`
	Coin       string             `bson:"coin,omitempty" json:"coin,omitempty"`
	Coins      string             `bson:"coins,omitempty" json:"coins,omitempty"`
	Address    string             `bson:"address,omitempty" json:"address,omitempty"`
	City       string             `bson:"city,omitempty" json:"city,omitempty"`
	ProvEstate string             `bson:"provEstate,omitempty" json:"provEstate,omitempty"`
	CreateDate *time.Time         `bson:"createDate,omitempty" json:"createDate,omitempty"`
	ExpireDate *time.Time         `bson:"expireDate,omitempty" json:"expireDate,omitempty"`
	Type       int                `bson:"type" json:"type"`
}
