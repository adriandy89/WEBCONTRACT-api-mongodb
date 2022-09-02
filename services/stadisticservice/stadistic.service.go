package stadisticservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/supplementservice"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TotalContractByCodeCompanyQueryClasif(code string) (int64, int64, int64, int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": code, "state": "Vigente", "clientSupplier": "Cliente"}
	condition2 := bson.M{"codeCompany": code, "state": "Terminado", "clientSupplier": "Cliente"}

	condition3 := bson.M{"codeCompany": code, "state": "Vigente", "clientSupplier": "Proveedor"}
	condition4 := bson.M{"codeCompany": code, "state": "Terminado", "clientSupplier": "Proveedor"}

	activeClient, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0, 0, 0, 0
	}
	inactiveClient, err2 := db.ContractCollection.CountDocuments(ctx, condition2)
	if err2 != nil {
		return 0, 0, 0, 0
	}

	activeProvider, err3 := db.ContractCollection.CountDocuments(ctx, condition3)
	if err3 != nil {
		return 0, 0, 0, 0
	}
	inactiveProvider, err4 := db.ContractCollection.CountDocuments(ctx, condition4)
	if err4 != nil {
		return 0, 0, 0, 0
	}

	return activeClient, inactiveClient, activeProvider, inactiveProvider

}

func FindActivesByCodeCompanyAndDate(codeCompany string) ([]*models.Contract, []*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	//clientes
	conditionClient := bson.M{"codeCompany": codeCompany, "state": "Vigente", "clientSupplier": "Cliente", "expireAt": bson.M{"$lt": time.Now().Add(-24 * time.Hour)}}
	var contractsClient []*models.Contract
	//proveedores
	conditionProv := bson.M{"codeCompany": codeCompany, "state": "Vigente", "clientSupplier": "Proveedor", "expireAt": bson.M{"$lt": time.Now().Add(-24 * time.Hour)}}
	var contractsProv []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, conditionClient)
	cursorProv, errProv := db.ContractCollection.Find(ctx, conditionProv)

	if err != nil {
		return nil, nil, false
	}
	err = cursor.Err()
	if err != nil {
		return nil, nil, false
	}
	if errProv != nil {
		return nil, nil, false
	}
	errProv = cursorProv.Err()
	if errProv != nil {
		return nil, nil, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contractClient models.Contract
		err := cursor.Decode(&contractClient)
		if err != nil {
			return nil, nil, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contractsClient = append(contractsClient, &contractClient)
	}

	defer cursorProv.Close(context.Background())
	for cursorProv.Next(context.Background()) {
		var contractProv models.Contract
		errProv := cursorProv.Decode(&contractProv)
		if errProv != nil {
			return nil, nil, false
		}
		contractsProv = append(contractsProv, &contractProv)
	}
	return contractsClient, contractsProv, true
}

func TotalTypeCoisByCodeCompany(code string) (int64, int64, int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{
		"codeCompany": code,
		"$or": []bson.M{
			bson.M{"codeTypeCoin": "CUC"},
			bson.M{"codeTypeCoin": "MN"},
			bson.M{"codeTypeCoin": "AMBAS"},
		},
	}
	condition2 := bson.M{"codeCompany": code, "codeTypeCoin": "MLC"}
	// Los contratos solo se pueden hacer en una moneda, por eso se kita est condicion
	//condition3 := bson.M{"codeCompany": code, "codeTypeCoin": "AMBAS"}

	cup, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0, 0, 0
	}
	mlc, err2 := db.ContractCollection.CountDocuments(ctx, condition2)
	if err2 != nil {
		return 0, 0, 0
	}
	/*ambas, err3 := db.ContractCollection.CountDocuments(ctx, condition3)
	if err3 != nil {
		return 0, 0, 0
	}*/

	return cup, mlc, 0
}

func FindActivesByCodeCompanyGroupBy(codeCompany string) ([]*models.Contract, []*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "state": "Vigente", "clientSupplier": "Cliente"}
	var contractsC []*models.Contract

	condition2 := bson.M{"codeCompany": codeCompany, "state": "Vigente", "clientSupplier": "Proveedor"}
	var contractsP []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"clientProviderName": 1}).SetProjection(bson.M{"clientProviderName": 1, "codeCategoryInitial": 1, "clientSupplier": 1}))
	cursor2, err2 := db.ContractCollection.Find(ctx, condition2, options.Find().SetSort(bson.M{"clientProviderName": 1}).SetProjection(bson.M{"clientProviderName": 1, "codeCategoryInitial": 1, "clientSupplier": 1}))

	if err != nil || err2 != nil {
		return contractsC, contractsP, false
	}
	err = cursor.Err()
	err2 = cursor2.Err()
	if err != nil || err2 != nil {
		return contractsC, contractsP, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.Contract
		err := cursor.Decode(&contract)
		if err != nil {
			return contractsC, contractsP, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contractsC = append(contractsC, &contract)
	}
	defer cursor2.Close(context.Background())
	for cursor2.Next(context.Background()) {
		var contract2 models.Contract
		err := cursor2.Decode(&contract2)
		if err != nil {
			return contractsC, contractsP, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contractsP = append(contractsP, &contract2)
	}

	return contractsC, contractsP, true
}

//------------ Range od Dates

func TotalTypeCoisByCodeCompanyXDate(code string, start *time.Time, end *time.Time) (int64, int64, int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{
		"codeCompany": code,
		"createdAt":   bson.M{"$gte": start, "$lt": end},
		"$or": []bson.M{
			bson.M{"codeTypeCoin": "CUC"},
			bson.M{"codeTypeCoin": "MN"},
			bson.M{"codeTypeCoin": "AMBAS"},
		},
	}
	condition2 := bson.M{"codeCompany": code, "codeTypeCoin": "MLC", "createdAt": bson.M{"$gte": start, "$lt": end}}

	//condition3 := bson.M{"codeCompany": code, "codeTypeCoin": "AMBAS", "createdAt": bson.M{"$gte": start, "$lt": end}}

	cup, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0, 0, 0
	}
	mlc, err2 := db.ContractCollection.CountDocuments(ctx, condition2)
	if err2 != nil {
		return 0, 0, 0
	}
	/* ambas, err3 := db.ContractCollection.CountDocuments(ctx, condition3)
	if err3 != nil {
		return 0, 0, 0
	} */

	return cup, mlc, 0
}

func TotalContractsByMonths(code string, start *time.Time, end *time.Time) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{
		"codeCompany": code,
		"createdAt":   bson.M{"$gte": start, "$lt": end},
	}

	cant, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0
	}

	return cant
}

func TotalSuplementsByMonths(code string, start *time.Time, end *time.Time) int64 {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{
		"codeCompany": code,
		"createdAt":   bson.M{"$gte": start, "$lt": end},
	}

	cant, err := db.SuplementCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0
	}
	return cant
}

func ContractsWithSupplTotalCoins(codeCompany string, typeContract string) ([]*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "clientSupplier": typeContract}
	var contracts []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetProjection(bson.M{"codeContract": 1, "codeReeup": 1, "codeCompany": 1, "ammountMN": 1, "ammountCUC": 1, "ammountMLC": 1, "createdAt": 1}))

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
		contract.Supplements, _ = supplementservice.FindAllByCodeCompanyContractReeup(contract.CodeCompany, contract.CodeContract, contract.CodeReeup)

		contracts = append(contracts, &contract)
	}
	return contracts, true
}
