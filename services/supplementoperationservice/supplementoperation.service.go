package supplementoperationservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllSupplementsOperations
func FindAllSupplementsOperations() ([]*models.SupplementOperation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var supplementoperations []*models.SupplementOperation
	cursor, err := db.SupplementsOperationsCollection.Find(ctx, bson.M{})
	if err != nil {
		return supplementoperations, err
	}
	err = cursor.Err()
	if err != nil {
		return supplementoperations, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var supplementoperation models.SupplementOperation
		err := cursor.Decode(&supplementoperation)
		if err != nil {
			return supplementoperations, err
		}
		supplementoperations = append(supplementoperations, &supplementoperation)
	}
	return supplementoperations, nil
}

// FindByID
func FindByID(id string) (models.SupplementOperation, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.SupplementOperation

	err := db.SupplementsOperationsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByDescriptionAndCode  ---- DescriptionQuery() y CodeQuery()--- concurrentes
func ValidateIfExistByDescriptionAndCode(description string, code int) bool {

	c := make(chan bool)
	d := make(chan bool)

	go DescriptionQuery(description, c)
	go CodeQuery(code, d)

	result := false

	if <-c || <-d {
		result = true
	}

	return result
}
func DescriptionQuery(description string, c chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"description": description}
	var result models.SupplementOperation

	err := db.SupplementsOperationsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code int, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeOperation": code}
	var result models.SupplementOperation

	err := db.SupplementsOperationsCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewSupplementOperation
func InsertNewSupplementOperation(c models.SupplementOperation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.SupplementsOperationsCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID
func UpdateByID(id string, cUpdate models.SupplementOperation) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeOperation": cUpdate.CodeOperation,
		"description":   cUpdate.Description,
	}}

	upd, err := db.SupplementsOperationsCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.SupplementsOperationsCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
