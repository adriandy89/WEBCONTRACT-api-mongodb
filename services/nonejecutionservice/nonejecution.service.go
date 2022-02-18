package nonejecutionservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllNonEjecutions => Devuelve todas los motivos de No Ejecucion
func FindAllNonEjecutions() ([]*models.NonEjecution, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var nonEjecutions []*models.NonEjecution
	cursor, err := db.NonEjecutionsCollection.Find(ctx, bson.M{})
	if err != nil {
		return nonEjecutions, err
	}
	err = cursor.Err()
	if err != nil {
		return nonEjecutions, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var nonEjecution models.NonEjecution
		err := cursor.Decode(&nonEjecution)
		if err != nil {
			return nonEjecutions, err
		}
		nonEjecutions = append(nonEjecutions, &nonEjecution)
	}
	return nonEjecutions, nil
}

// FindByID => Funcion para obtener entidad por id
func FindByID(id string) (models.NonEjecution, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.NonEjecution

	err := db.NonEjecutionsCollection.FindOne(ctx, condition).Decode(&result)
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
	var result models.NonEjecution

	err := db.NonEjecutionsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code int, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeNonExecution": code}
	var result models.NonEjecution

	err := db.NonEjecutionsCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewNonEjecution
func InsertNewNonEjecution(c models.NonEjecution) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.NonEjecutionsCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID
func UpdateByID(id string, cUpdate models.NonEjecution) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeNonExecution": cUpdate.CodeNonExecution,
		"description":      cUpdate.Description,
	}}

	upd, err := db.NonEjecutionsCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.NonEjecutionsCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
