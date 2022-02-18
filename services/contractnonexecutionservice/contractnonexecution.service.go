package contractnonexecutionservice

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
func FindByCountAndSort(codeCompany string, count int, order string, typ string, page int) ([]*models.ContractNonExecution, int64, bool) {

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

	var contracts []*models.ContractNonExecution
	cursor, err := db.ContractNonExecutionCollection.Find(ctx, condition, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
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
		var contract models.ContractNonExecution
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

	cursor, err := db.ContractNonExecutionCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

// FindByID => Funcion para obtener contratos por id
func FindByID(id string) (models.ContractNonExecution, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.ContractNonExecution

	err := db.ContractNonExecutionCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// InsertNewContract => guarda un nuevo contrato en la base de datos
func InsertNewContract(u models.ContractNonExecution) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.ContractNonExecutionCollection.InsertOne(ctx, u)

	return err
}

// DeleteByID => Funcion para eliminar contrato por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.ContractNonExecutionCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0
}

// UpdateByID => Funcion para actualizar contrato por id
func UpdateByID(id string, cUpdate models.ContractNonExecution) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeContract":           cUpdate.CodeContract,
		"codeReeup":              cUpdate.CodeReeup,
		"codeCompany":            cUpdate.CodeCompany,
		"idNonExecution":         cUpdate.IdNonExecution,
		"contractNonExecution":   cUpdate.ContractNonExecution,
		"idContractNonExecution": cUpdate.IdContractNonExecution,
	}}

	upd, err := db.ContractNonExecutionCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}
