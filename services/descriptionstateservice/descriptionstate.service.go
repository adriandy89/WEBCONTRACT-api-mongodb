package descriptionstateservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindAllDescriptionsStates => Devuelve todas las Descipciones de estado
func FindAllDescriptionsStates() ([]*models.DescriptionState, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var descriptionsStates []*models.DescriptionState
	cursor, err := db.DescriptionsStatesCollection.Find(ctx, bson.M{}, options.Find().SetSort(bson.M{"description": 1}))
	if err != nil {
		return descriptionsStates, err
	}
	err = cursor.Err()
	if err != nil {
		return descriptionsStates, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var descriptionState models.DescriptionState
		err := cursor.Decode(&descriptionState)
		if err != nil {
			return descriptionsStates, err
		}
		descriptionsStates = append(descriptionsStates, &descriptionState)
	}
	return descriptionsStates, nil
}

// FindByID => Funcion para obtener Descipcion de estado por id
func FindByID(id string) (models.DescriptionState, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.DescriptionState

	err := db.DescriptionsStatesCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByDescriptionAndCode => validar si existe la Descipcion de estado ---- NameQuery() y CodeQuery()--- concurrentes
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
	var result models.DescriptionState

	err := db.DescriptionsStatesCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code int, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeDescriptionState": code}
	var result models.DescriptionState

	err := db.DescriptionsStatesCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewCurrency => Inserta nueva Descipcion de estado en la base de datos
func InsertNewDescriptionState(c models.DescriptionState) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.DescriptionsStatesCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID => Funcion para actualizar Descipcion de estado por id
func UpdateByID(id string, cUpdate models.DescriptionState) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeDescriptionState": cUpdate.CodeDescriptionState,
		"description":          cUpdate.Description,
	}}

	upd, err := db.DescriptionsStatesCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID => Funcion para eliminar Descipcion de estado por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.DescriptionsStatesCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
