package stadisticservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TotalContractByCodeCompanyQueryClasif(code string) (int64, int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": code, "state": "Vigente"}
	condition2 := bson.M{"codeCompany": code, "state": "Terminado"}

	active, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0, 0
	}
	inactive, err2 := db.ContractCollection.CountDocuments(ctx, condition2)
	if err2 != nil {
		return 0, 0
	}

	return active, inactive

}

func FindActivesByCodeCompanyAndDate(codeCompany string) ([]*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "state": "Vigente", "expireAt": bson.M{"$lt": time.Now().Add(-24 * time.Hour)}}
	var contracts []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, condition)

	if err != nil {
		return contracts, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.Contract
		err := cursor.Decode(&contract)
		if err != nil {
			return contracts, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contracts = append(contracts, &contract)
	}
	return contracts, true
}

func TotalTypeCoisByCodeCompany(code string) (int64, int64, int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{
		"codeCompany": code,
		"$or": []bson.M{
			bson.M{"codeTypeCoin": "CUC"},
			bson.M{"codeTypeCoin": "MN"},
		},
	}
	condition2 := bson.M{"codeCompany": code, "codeTypeCoin": "MLC"}
	condition3 := bson.M{"codeCompany": code, "codeTypeCoin": "AMBAS"}

	cup, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0, 0, 0
	}
	mlc, err2 := db.ContractCollection.CountDocuments(ctx, condition2)
	if err2 != nil {
		return 0, 0, 0
	}
	ambas, err3 := db.ContractCollection.CountDocuments(ctx, condition3)
	if err3 != nil {
		return 0, 0, 0
	}

	return cup, mlc, ambas
}

func FindActivesByCodeCompanyGroupBy(codeCompany string) ([]*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "state": "Vigente"}
	var contracts []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"clientProviderName": 1}).SetProjection(bson.M{"clientProviderName": 1}))

	if err != nil {
		return contracts, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.Contract
		err := cursor.Decode(&contract)
		if err != nil {
			return contracts, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contracts = append(contracts, &contract)
	}
	return contracts, true
}
