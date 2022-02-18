package currencyservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindAllCurrencies => Devuelve todas las Monedas
func FindAllCurrencies() ([]*models.Currency, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var currencies []*models.Currency
	cursor, err := db.CurrencyCollection.Find(ctx, bson.M{})
	if err != nil {
		return currencies, err
	}
	err = cursor.Err()
	if err != nil {
		return currencies, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var currency models.Currency
		err := cursor.Decode(&currency)
		if err != nil {
			return currencies, err
		}
		currencies = append(currencies, &currency)
	}
	return currencies, nil
}

// FindByID => Funcion para obtener moneda por id
func FindByID(id string) (models.Currency, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.Currency

	err := db.CurrencyCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByNameAndCode => validar si existe la moneda ---- NameQuery() y CodeQuery()--- concurrentes
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
	var result models.Currency

	err := db.CurrencyCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"currency": code}
	var result models.Currency

	err := db.CurrencyCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewCurrency => Inserta nueva Moneda en la base de datos
func InsertNewCurrency(c models.Currency) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.CurrencyCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID => Funcion para actualizar moneda por id
func UpdateByID(id string, cUpdate models.Currency) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"currency": cUpdate.Currency,
		"name":     cUpdate.Name,
	}}

	upd, err := db.CurrencyCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID => Funcion para eliminar moneda por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.CurrencyCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
