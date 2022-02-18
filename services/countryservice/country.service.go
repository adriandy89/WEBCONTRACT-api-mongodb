package countryservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllCountries => Devuelve todas las provincias
func FindAllCountries() ([]*models.Country, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var countries []*models.Country
	cursor, err := db.CountryCollection.Find(ctx, bson.M{})
	if err != nil {
		return countries, err
	}
	err = cursor.Err()
	if err != nil {
		return countries, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var country models.Country
		err := cursor.Decode(&country)
		if err != nil {
			return countries, err
		}
		countries = append(countries, &country)
	}
	return countries, nil
}

// FindByID => Funcion para obtener provincia por id
func FindByID(id string) (models.Country, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.Country

	err := db.CountryCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByNameAndCode => validar si existe la provincia ---- NameQuery() y CodeQuery()--- concurrentes
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
	var result models.Country

	err := db.CountryCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeCountry": code}
	var result models.Country

	err := db.CountryCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewCountry => Inserta nueva Provincia en la base de datos
func InsertNewCountry(c models.Country) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.CountryCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID => Funcion para actualizar categoria por id
func UpdateByID(id string, cUpdate models.Country) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeCountry": cUpdate.CodeCountry,
		"name":        cUpdate.Name,
		"initial":     cUpdate.Initial,
	}}

	upd, err := db.CountryCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID => Funcion para eliminar categoria por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.CountryCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
