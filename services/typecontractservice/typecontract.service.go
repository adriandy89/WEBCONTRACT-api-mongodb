package typecontractservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllTypeContracts
func FindAllTypeContracts() ([]*models.TypeContract, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var typeContracts []*models.TypeContract
	cursor, err := db.TypeContractsCollection.Find(ctx, bson.M{})
	if err != nil {
		return typeContracts, err
	}
	err = cursor.Err()
	if err != nil {
		return typeContracts, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var typeContract models.TypeContract
		err := cursor.Decode(&typeContract)
		if err != nil {
			return typeContracts, err
		}
		typeContracts = append(typeContracts, &typeContract)
	}
	return typeContracts, nil
}

// FindByID
func FindByID(id string) (models.TypeContract, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.TypeContract

	err := db.TypeContractsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByNameAndCode  ---- DescriptionQuery() y CodeQuery()--- concurrentes
func ValidateIfExistByNameAndCode(name string, code string) bool {

	c := make(chan bool)
	d := make(chan bool)

	go NameQuery(name, c)
	go CodeQuery(code, d)

	result := false

	if <-c || <-d {
		result = true
	}

	return result
}
func NameQuery(name string, c chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"name": name}
	var result models.TypeContract

	err := db.TypeContractsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeTypeContract": code}
	var result models.TypeContract

	err := db.TypeContractsCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewTypeContract
func InsertNewTypeContract(c models.TypeContract) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.TypeContractsCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID
func UpdateByID(id string, cUpdate models.TypeContract) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeTypeContract": cUpdate.CodeTypeContract,
		"name":             cUpdate.Name,
	}}

	upd, err := db.TypeContractsCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.TypeContractsCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
