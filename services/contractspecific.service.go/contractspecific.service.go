package contractspecificservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindByCount => devuelve los Contratos filtrados  ---------- TotalContractQuery() --- concurrentes
func FindByCountAndSort(codeCompany string, count int, order string, typ string, page int) ([]*models.ContractSpecific, int64, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	c := make(chan int64)

	var sort int = 1
	if typ == "d" {
		sort = -1
	}
	pageNumber := 0
	if page > 1 {
		pageNumber = (page - 1) * count
	}

	condition := bson.M{"codeCompany": codeCompany}

	var contracts []*models.ContractSpecific
	cursor, err := db.ContractSpecificCollection.Find(ctx, condition, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return contracts, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, 0, false
	}

	go TotalContractQuery(c)

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.ContractSpecific
		err := cursor.Decode(&contract)
		if err != nil {
			return contracts, 0, false
		}
		contracts = append(contracts, &contract)
	}
	return contracts, <-c, true
}
func TotalContractQuery(c chan int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := db.ContractSpecificCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

// ValidateIfExistByCodeCompanyCodeContractAndCodeSpecific
func ValidateIfExistByCodeCompanyCodeContractAndCodeSpecific(codeCompany string, codeContract string, codeSpecific string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "codeContract": codeContract, "codeSpecific": codeSpecific}

	var result models.ContractSpecific

	err := db.ContractSpecificCollection.FindOne(ctx, condition).Decode(&result)

	return err == nil
}

// ValidateIfExistByCodeCompanyAndCodeContract
func ValidateIfExistByCodeCompanyAndCodeContract(codeCompany string, codeContract string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "codeContract": codeContract}

	var result models.ContractSpecific

	err := db.ContractSpecificCollection.FindOne(ctx, condition).Decode(&result)

	return err == nil
}

// InsertNewContract => guarda un nuevo contrato en la base de datos
func InsertNewContractSpecific(u models.ContractSpecific) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.ContractSpecificCollection.InsertOne(ctx, u)

	return err
}

// DeleteByID => Funcion para eliminar contrato por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.ContractSpecificCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0
}

// FindByID => Funcion para obtener contratos por id
func FindByID(id string) (models.ContractSpecific, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.ContractSpecific

	err := db.ContractSpecificCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// UpdateByID => Funcion para actualizar contrato por id
func UpdateByID(id string, cUpdate models.ContractSpecific) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeSpecific":            cUpdate.CodeSpecific,
		"codeContract":            cUpdate.CodeContract,
		"codeReeup":               cUpdate.CodeReeup,
		"codeCompany":             cUpdate.CodeCompany,
		"createdAt":               cUpdate.CreatedAt,
		"expireAt":                cUpdate.ExpireAt,
		"codeTypeCoin":            cUpdate.CodeTypeCoin,
		"codeTypeContract":        cUpdate.CodeTypeContract,
		"state":                   cUpdate.State,
		"ammountMN":               cUpdate.AmmountMN,
		"ammountCUC":              cUpdate.AmmountCUC,
		"verdictLegal":            cUpdate.VerdictLegal,
		"processPersonI":          cUpdate.ProcessPersonI,
		"processPersonF":          cUpdate.ProcessPersonF,
		"nroArchive":              cUpdate.NroArchive,
		"payPeriod":               cUpdate.PayPeriod,
		"creditDays":              cUpdate.CreditDays,
		"clientSupplier":          cUpdate.ClientSupplier,
		"codeOfert":               cUpdate.CodeOfert,
		"codeObject":              cUpdate.CodeObject,
		"ammountMNSuplementary":   cUpdate.AmmountMNSuplementary,
		"ammountCUCSuplementary":  cUpdate.AmmountCUCSuplementary,
		"codeDescriptionState":    cUpdate.CodeDescriptionState,
		"commentDescriptionState": cUpdate.CommentDescriptionState,
		"paymentTerm":             cUpdate.PaymentTerm,
		"codeCategory":            cUpdate.CodeCategory,
		"ammountMNInit":           cUpdate.AmmountMNInit,
		"fileRoute":               cUpdate.FileRoute,
	}}

	upd, err := db.ContractSpecificCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

func CodeSpecificQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeSpecific": code}
	var result models.ContractSpecific

	err := db.ContractSpecificCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// GetNewCodeContractSpecific
func GetNewCodeContractSpecific(codeCompany string, codeContract string) ([]string, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "codeContract": codeContract}
	condition2 := bson.M{"codeSpecific": 1}

	var list []string
	cursor, err := db.ContractSpecificCollection.Find(ctx, condition, options.Find().SetProjection(condition2))
	if err != nil {
		return list, false
	}
	err = cursor.Err()
	if err != nil {
		return list, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.ContractSpecific
		err := cursor.Decode(&contract)
		if err != nil {
			return list, false
		}
		list = append(list, contract.CodeSpecific)
	}
	return list, true
}
