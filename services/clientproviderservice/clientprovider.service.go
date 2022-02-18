package clientproviderservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindByCount => devuelve los Clientes y proveedores filtrados  ---------- TotalClientProviderQuery() --- concurrentes
func FindByCountAndSort(count int, order string, typ string, page int) ([]*models.ClientProvider, int64, bool) {

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

	var clientProviders []*models.ClientProvider
	cursor, err := db.ClientProviderCollection.Find(ctx, bson.M{}, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return clientProviders, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return clientProviders, 0, false
	}

	go TotalClientProviderQuery(c)

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var clientProvider models.ClientProvider
		err := cursor.Decode(&clientProvider)
		if err != nil {
			return clientProviders, 0, false
		}
		clientProviders = append(clientProviders, &clientProvider)
	}
	return clientProviders, <-c, true
}
func TotalClientProviderQuery(c chan int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := db.ClientProviderCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

// ValidateIfExistByCustId => buscar un Cliente o Provedor en la base de datos por custId
func ValidateIfExistByCustId(custId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"custId": custId}

	var result models.ClientProvider

	err := db.ClientProviderCollection.FindOne(ctx, condition).Decode(&result)

	return err == nil
}

// InsertNewClientProvider => guarda un nuevo cliente o provedor en la base de datos
func InsertNewClientProvider(u models.ClientProvider) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.ClientProviderCollection.InsertOne(ctx, u)

	return err
}

// DeleteByID => Funcion para eliminar categoria por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cpID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": cpID}

	delRes, err := db.ClientProviderCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0
}

// FindByID => Funcion para obtener categoria por id
func FindByID(id string) (models.ClientProvider, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cpID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": cpID}
	var result models.ClientProvider

	err := db.ClientProviderCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// FindNameByCustID => Funcion para obtener categoria por custId
func FindNameByCustID(id string) (string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"custId": id}
	var result models.ClientProvider

	err := db.ClientProviderCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result.Name, false
	}
	return result.Name, true
}

// UpdateByID => Funcion para actualizar cliente o proveedor por id
func UpdateByID(id string, cpUpdate models.ClientProvider) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cpID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": cpID}

	update := bson.M{"$set": bson.M{
		"custId":     cpUpdate.CustId,
		"codOne":     cpUpdate.CodOne,
		"name":       cpUpdate.Name,
		"organism":   cpUpdate.Organism,
		"statusCode": cpUpdate.StatusCode,
		"email":      cpUpdate.Email,
		"phone":      cpUpdate.Phone,
		"country":    cpUpdate.Country,
		"sectorType": cpUpdate.SectorType,
		"coin":       cpUpdate.Coin,
		"coins":      cpUpdate.Coins,
		"address":    cpUpdate.Address,
		"city":       cpUpdate.City,
		"provEstate": cpUpdate.ProvEstate,
		"createDate": cpUpdate.CreateDate,
		"expireDate": cpUpdate.ExpireDate,
		"type":       cpUpdate.Type,
	}}

	upd, err := db.ClientProviderCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

func CustIdQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"custId": code}
	var result models.ClientProvider

	err := db.ClientProviderCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// FindByNameOrCode => devuelve los Clientes y proveedores filtrados  ---------- TotalClientProviderQuery() --- concurrentes
func FindByNameOrCode(count int, order string, typ string, page int, word string) ([]*models.ClientProvider, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var sort int = 1
	if typ == "d" {
		sort = -1
	}
	pageNumber := 0
	if page > 1 {
		pageNumber = (page - 1) * count
	}

	var clientProviders []*models.ClientProvider
	cursor, err := db.ClientProviderCollection.Find(ctx,
		bson.M{
			"$or": []bson.M{
				bson.M{"name": bson.M{"$regex": word, "$options": "im"}},
				bson.M{"custId": bson.M{"$regex": word, "$options": "im"}},
			},
		},
		options.Find().SetLimit(int64(count)),
		options.Find().SetProjection(bson.D{{"name", 1}, {"custId", 1}, {"type", 1}}),
		options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return clientProviders, false
	}
	err = cursor.Err()
	if err != nil {
		return clientProviders, false
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var clientProvider models.ClientProvider
		err := cursor.Decode(&clientProvider)
		if err != nil {
			return clientProviders, false
		}
		clientProviders = append(clientProviders, &clientProvider)
	}
	return clientProviders, true
}
