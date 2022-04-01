package models

import "time"

// FindCodeContractSpecific => Estructura de la solicitud de un nuevo codigo
type FindCodeContractSpecific struct {
	CodeContract string `bson:"codeContract" json:"codeContract"`
	CodeCompany  string `bson:"codeCompany" json:"codeCompany"`
}

type Word struct {
	Word string `bson:"word" json:"word"`
}

type At struct {
	At *time.Time `bson:"at" json:"at"`
}
