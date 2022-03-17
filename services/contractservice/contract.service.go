package contractservice

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
func FindByCountAndSort(codeCompany string, count int, order string, typ string, page int) ([]*models.Contract, int64, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	c := make(chan int64)

	var sort int = 1
	if typ == "desc" {
		sort = -1
	}
	pageNumber := 0
	if page > 1 {
		pageNumber = (page - 1) * count
	}

	condition := bson.M{"codeCompany": codeCompany}

	var contracts []*models.Contract
	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return contracts, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, 0, false
	}

	go TotalContractByCodeCompanyQuery(c, codeCompany)

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.Contract
		err := cursor.Decode(&contract)
		if err != nil {
			return contracts, 0, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contracts = append(contracts, &contract)
	}
	return contracts, <-c, true
}

func FindByNameOrCode(codeCompany string, count int, order string, typ string, page int, word string) ([]*models.Contract, int64, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	c := make(chan int64)
	var sort int = 1
	if typ == "desc" {
		sort = -1
	}
	pageNumber := 0
	if page > 1 {
		pageNumber = (page - 1) * count
	}
	go TotalContractsQueryByWord(c, word, codeCompany)
	var contracts []*models.Contract
	cursor, err := db.ContractCollection.Find(ctx,
		bson.M{
			"codeCompany": codeCompany,
			"$or": []bson.M{
				bson.M{"clientProviderName": bson.M{"$regex": word, "$options": "im"}},
				bson.M{"codeContract": bson.M{"$regex": word, "$options": "im"}},
			},
		},
		options.Find().SetLimit(int64(count)),
		options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return contracts, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, 0, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.Contract
		err := cursor.Decode(&contract)
		if err != nil {
			return contracts, 0, false
		}
		contracts = append(contracts, &contract)
	}
	return contracts, <-c, true
}

func TotalContractsQueryByWord(c chan int64, word string, codeCompany string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := db.ContractCollection.CountDocuments(ctx, bson.M{
		"codeCompany": codeCompany,
		"$or": []bson.M{
			bson.M{"clientProviderName": bson.M{"$regex": word, "$options": "im"}},
			bson.M{"codeContract": bson.M{"$regex": word, "$options": "im"}},
		},
	})
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

func TotalContractByCodeCompanyQuery(c chan int64, code string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": code}

	cursor, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

// ValidateIfExistByCodeContract => buscar un Contrato en la base de datos por codeContract
func ValidateIfExistByCodeContract(code string, codeCompany string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeContract": code, "codeCompany": codeCompany}

	var result models.Contract

	err := db.ContractCollection.FindOne(ctx, condition).Decode(&result)

	return err == nil
}

// InsertNewContract => guarda un nuevo contrato en la base de datos
func InsertNewContract(u models.Contract) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.ContractCollection.InsertOne(ctx, u)

	return err
}

// DeleteByID => Funcion para eliminar contrato por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.ContractCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0
}

// FindByID => Funcion para obtener contratos por id
func FindByID(id string) (models.Contract, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.Contract

	err := db.ContractCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// UpdateByID => Funcion para actualizar contrato por id
func UpdateByID(id string, cUpdate models.Contract) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
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
		"supplements":             cUpdate.Supplements,
		"nonCompliance":           cUpdate.NonCompliance,
	}}

	upd, err := db.ContractCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

func CodeContractQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeContract": code}
	var result models.Contract

	err := db.ContractCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// GetNewCodeContract => devuelve una Lista con con posibles nuevos codigos de Contratos
func GetNewCodeContract(codeCompany string, year string) ([]string, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "codeContract": bson.M{"$regex": primitive.Regex{Pattern: year + "$", Options: "i"}}}
	condition2 := bson.M{"codeContract": 1}

	var list []string
	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetProjection(condition2))
	if err != nil {
		return list, false
	}
	err = cursor.Err()
	if err != nil {
		return list, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.Contract
		err := cursor.Decode(&contract)
		if err != nil {
			return list, false
		}
		list = append(list, contract.CodeContract)
	}
	return list, true
}
