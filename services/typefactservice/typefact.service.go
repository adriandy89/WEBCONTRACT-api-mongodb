package typefactservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllTypeFacts
func FindAllTypeFacts() ([]*models.TypeFact, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var typeFacts []*models.TypeFact
	cursor, err := db.TypeFactsCollection.Find(ctx, bson.M{})
	if err != nil {
		return typeFacts, err
	}
	err = cursor.Err()
	if err != nil {
		return typeFacts, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var typeFact models.TypeFact
		err := cursor.Decode(&typeFact)
		if err != nil {
			return typeFacts, err
		}
		typeFacts = append(typeFacts, &typeFact)
	}
	return typeFacts, nil
}

// FindByID
func FindByID(id string) (models.TypeFact, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.TypeFact

	err := db.TypeFactsCollection.FindOne(ctx, condition).Decode(&result)
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
	var result models.TypeFact

	err := db.TypeFactsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code int, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeFact": code}
	var result models.TypeFact

	err := db.TypeFactsCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewTypeFact
func InsertNewTypeFact(c models.TypeFact) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.TypeFactsCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID
func UpdateByID(id string, cUpdate models.TypeFact) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeFact":    cUpdate.CodeFact,
		"description": cUpdate.Description,
	}}

	upd, err := db.TypeFactsCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.TypeFactsCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
