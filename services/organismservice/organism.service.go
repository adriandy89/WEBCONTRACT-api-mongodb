package organismservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllOrganisms => Devuelve todas los organismos
func FindAllOrganisms() ([]*models.Organism, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var organisms []*models.Organism
	cursor, err := db.OrganismsCollection.Find(ctx, bson.M{})
	if err != nil {
		return organisms, err
	}
	err = cursor.Err()
	if err != nil {
		return organisms, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var organism models.Organism
		err := cursor.Decode(&organism)
		if err != nil {
			return organisms, err
		}
		organisms = append(organisms, &organism)
	}
	return organisms, nil
}

// FindByID
func FindByID(id string) (models.Organism, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.Organism

	err := db.OrganismsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByNameAndCode ---- NameQuery() y CodeQuery()--- concurrentes
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
	var result models.Organism

	err := db.OrganismsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeOrganism": code}
	var result models.Organism

	err := db.OrganismsCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewOrganism => Inserta nueva Provincia en la base de datos
func InsertNewOrganism(c models.Organism) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.OrganismsCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID
func UpdateByID(id string, cUpdate models.Organism) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeOrganism": cUpdate.CodeOrganism,
		"name":         cUpdate.Name,
		"initial":      cUpdate.Initial,
	}}

	upd, err := db.OrganismsCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.OrganismsCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
