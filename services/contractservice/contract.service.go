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
func FindByCountAndSort(codeCompany string, count int, order string, typ string, page int, state string) ([]*models.Contract, int64, bool) {

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
	condition := bson.M{}

	if state == "" {
		condition = bson.M{"codeCompany": codeCompany}
	} else if state == "Vigente" {
		condition = bson.M{"codeCompany": codeCompany, "state": state}
	}

	var contracts []*models.Contract
	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return contracts, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, 0, false
	}

	if state == "Vigente" {
		go TotalContractByCodeCompanyVigentQuery(c, codeCompany)
	} else {
		go TotalContractByCodeCompanyQuery(c, codeCompany)
	}

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

func FindByNameOrCode(codeCompany string, count int, order string, typ string, page int, word, state string) ([]*models.Contract, int64, bool) {

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
	condition := bson.M{}

	if state == "" {
		condition = bson.M{
			"codeCompany": codeCompany,
			"$or": []bson.M{
				bson.M{"clientProviderName": bson.M{"$regex": word, "$options": "im"}},
				bson.M{"codeContract": bson.M{"$regex": word, "$options": "im"}},
			},
		}
	} else if state == "Vigente" {
		condition = bson.M{
			"codeCompany": codeCompany,
			"state":       state,
			"$or": []bson.M{
				bson.M{"clientProviderName": bson.M{"$regex": word, "$options": "im"}},
				bson.M{"codeContract": bson.M{"$regex": word, "$options": "im"}},
			},
		}
	}

	go TotalContractsQueryByWord(c, word, codeCompany, state)
	var contracts []*models.Contract
	cursor, err := db.ContractCollection.Find(ctx,
		condition,
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

func TotalContractsQueryByWord(c chan int64, word string, codeCompany string, state string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{}
	if state == "" {
		condition = bson.M{
			"codeCompany": codeCompany,
			"$or": []bson.M{
				bson.M{"clientProviderName": bson.M{"$regex": word, "$options": "im"}},
				bson.M{"codeContract": bson.M{"$regex": word, "$options": "im"}},
			},
		}
	} else if state == "Vigente" {
		condition = bson.M{
			"codeCompany": codeCompany,
			"state":       state,
			"$or": []bson.M{
				bson.M{"clientProviderName": bson.M{"$regex": word, "$options": "im"}},
				bson.M{"codeContract": bson.M{"$regex": word, "$options": "im"}},
			},
		}
	}

	cursor, err := db.ContractCollection.CountDocuments(ctx, condition)
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
func TotalContractByCodeCompanyVigentQuery(c chan int64, code string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": code, "state": "Vigente"}

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
		"ammountMLC":              cUpdate.AmmountMLC,
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
		"codeCategoryInitial":     cUpdate.CodeCategoryInitial,
		"ammountMNInit":           cUpdate.AmmountMNInit,
		"fileRoute":               cUpdate.FileRoute,
		"supplements":             cUpdate.Supplements,
		"nonCompliance":           cUpdate.NonCompliance,
		"cronogram":           	   cUpdate.Cronogram,
		"aproved":           	   cUpdate.Aproved,
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

// FindByCodeCompanyAndDate
func FindByCodeCompanyAndDate(codeCompany string) ([]*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "state": "Vigente", "expireAt": bson.M{"$lt": time.Now().Add(2880 * time.Hour)}}
	var contracts []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"clientProviderName": 1}))

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

// FindByCodeCompanyAndDate
func FindByCodeCompanyAndDateEXCEL(codeCompany string) ([]*models.ContractEXCEL, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "state": "Vigente", "expireAt": bson.M{"$lt": time.Now().Add(2880 * time.Hour)}}
	var contracts []*models.ContractEXCEL

	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"clientProviderName": 1}))

	if err != nil {
		return contracts, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.ContractEXCEL
		err := cursor.Decode(&contract)
		if err != nil {
			return contracts, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contracts = append(contracts, &contract)
	}
	return contracts, true
}

// FindByCodeCompanyAndDate   ---------- TotalContractQuery() --- concurrentes
func FindByCodeCompanyAndSpecificDate(codeCompany string, at *time.Time) ([]*models.Contract, int64, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	c := make(chan int64)

	condition := bson.M{"codeCompany": codeCompany, "createdAt": bson.M{"$gte": at}}
	var contracts []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, condition)

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

// GetDaysByDefaultContract
func GetDaysByDefaultContract(codeCompany string) int {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany}
	condition2 := bson.M{"range": 1}
	var rang models.Entity
	days := 0
	err := db.EntityCollection.FindOne(ctx, condition, options.FindOne().SetProjection(condition2)).Decode(&rang)
	if err == nil {
		days = rang.Range
	}
	return days
}

//
// Stadistics
//
// FindByCodeCompanyAndDateStadistic   ---------- TotalContractQuery() --- concurrentes
func FindByCodeCompanyAndDateStadistic(codeCompany string) ([]*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": codeCompany, "state": "Vigente", "expireAt": bson.M{"$lt": time.Now().Add(2880 * time.Hour)}}
	var contracts []*models.Contract

	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"clientProviderName": 1}))

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

func FindByCountAndSortStadistic(codeCompany string, count int, order string, typ string, page int, filter string) ([]*models.Contract, int64, bool) {

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
	condition := bson.M{}
	if filter == "Vencidos" {
		condition = bson.M{"codeCompany": codeCompany, "state": "Vigente", "expireAt": bson.M{"$lt": time.Now().Add(-24 * time.Hour)}}
	} else {
		condition = bson.M{"codeCompany": codeCompany, "state": filter}
	}

	var contracts []*models.Contract
	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return contracts, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, 0, false
	}

	go TotalContractByCodeCompanyQueryStadistic(c, condition)

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

func TotalContractByCodeCompanyQueryStadistic(c chan int64, condition primitive.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := db.ContractCollection.CountDocuments(ctx, condition)
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

func FindByCountAndSortStadisticAll(codeCompany string, filter string) ([]*models.Contract, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{}
	if filter == "Vencidos" {
		condition = bson.M{"codeCompany": codeCompany, "state": "Vigente", "expireAt": bson.M{"$lt": time.Now().Add(-24 * time.Hour)}}
	} else {
		condition = bson.M{"codeCompany": codeCompany, "state": "Vigente"}
	}

	var contracts []*models.Contract
	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"expireAt": 1}))
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

func FindByCountAndSortStadisticAllEXCEL(codeCompany string, filter string) ([]*models.ContractEXCEL, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{}
	if filter == "Terminado" {
		condition = bson.M{"codeCompany": codeCompany, "state": "Terminado"}
	} else if filter == "Vencidos" {
		condition = bson.M{"codeCompany": codeCompany, "state": "Vigente", "expireAt": bson.M{"$lt": time.Now().Add(-24 * time.Hour)}}
	} else {
		condition = bson.M{"codeCompany": codeCompany, "state": "Vigente"}
	}

	var contracts []*models.ContractEXCEL
	cursor, err := db.ContractCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"expireAt": 1}))
	if err != nil {
		return contracts, false
	}
	err = cursor.Err()
	if err != nil {
		return contracts, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var contract models.ContractEXCEL
		err := cursor.Decode(&contract)
		if err != nil {
			return contracts, false
		}
		//contract.ClientProviderName, _ = clientproviderservice.FindNameByCustID(contract.CodeReeup)
		contracts = append(contracts, &contract)
	}
	return contracts, true
}
