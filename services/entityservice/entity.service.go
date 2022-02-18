package entityservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllEntities => Devuelve todas las entidades
func FindAllEntities() ([]*models.Entity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var entities []*models.Entity
	cursor, err := db.EntityCollection.Find(ctx, bson.M{})
	if err != nil {
		return entities, err
	}
	err = cursor.Err()
	if err != nil {
		return entities, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var entity models.Entity
		err := cursor.Decode(&entity)
		if err != nil {
			return entities, err
		}
		entities = append(entities, &entity)
	}
	return entities, nil
}

// FindByID => Funcion para obtener entidad por id
func FindByID(id string) (models.Entity, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.Entity

	err := db.EntityCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByNameAndCode => validar si existe la Entidad ---- NameQuery() y CodeQuery()--- concurrentes
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

	condition := bson.M{"companyName": name}
	var result models.Entity

	err := db.EntityCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeCompany": code}
	var result models.Entity

	err := db.EntityCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewEntity => Inserta nueva Entidad en la base de datos
func InsertNewEntity(c models.Entity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.EntityCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID => Funcion para actualizar una entidad por id
func UpdateByID(id string, cUpdate models.Entity) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeCompany": cUpdate.CodeCompany,
		"address":     cUpdate.Address,
		"prefix":      cUpdate.Prefix,
		"companyName": cUpdate.CompanyName,
		"codeFather":  cUpdate.CodeFather,
	}}

	upd, err := db.EntityCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID => Funcion para eliminar una entidad por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.EntityCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}