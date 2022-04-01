package offerrequestservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"WEBCONTRACT-api-mongodb/services/clientproviderservice"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertNewOfferRequest
func InsertNewOfferRequest(u models.OfferRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.OffersRequestsCollection.InsertOne(ctx, u)

	return err
}

// ValidateIfExistByCodeOffer => buscar una Oferta en la base de datos por codeOffer
func ValidateIfExistByCodeOfferAndReeup(code string, codeCompany string, codeReeup string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeOffer": code, "codeCompany": codeCompany, "codeReeup": codeReeup}

	var result models.OfferRequest

	err := db.OffersRequestsCollection.FindOne(ctx, condition).Decode(&result)

	return err == nil
}

func CodeOfferQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeOffer": code}
	var result models.OfferRequest

	err := db.OffersRequestsCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// FindByID
func FindByID(id string) (models.OfferRequest, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.OfferRequest

	err := db.OffersRequestsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// FindAllByCompanyAndReeup
func FindAllByCompanyAndReeup(codeCompany string, codeReeup string) ([]*models.OfferRequest, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*models.OfferRequest

	condition := bson.M{"codeCompany": codeCompany, "codeReeup": codeReeup, "state": "Activo"}
	cursor, err := db.OffersRequestsCollection.Find(ctx, condition)
	if err != nil {
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var offer models.OfferRequest
		err := cursor.Decode(&offer)
		if err != nil {
			return results, false
		}
		results = append(results, &offer)
	}

	return results, true
}

// FindOneByCompanyReeupAndOffer
func FindOneByCompanyReeupAndOffer(codeCompany string, codeReeup string, codeOffer string) (*models.OfferRequest, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var result *models.OfferRequest

	condition := bson.M{"codeCompany": codeCompany, "codeReeup": codeReeup, "codeOffer": codeOffer}
	err := db.OffersRequestsCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// DeleteByID
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}

	delRes, err := db.OffersRequestsCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}

// UpdateByID
func UpdateByID(id string, cUpdate models.OfferRequest) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeReeup":          cUpdate.CodeReeup,
		"codeOffer":          cUpdate.CodeOffer,
		"ammountMN":          cUpdate.AmmountMN,
		"ammountCUC":         cUpdate.AmmountCUC,
		"state":              cUpdate.State,
		"description":        cUpdate.Description,
		"codeCompany":        cUpdate.CodeCompany,
		"fileRoute":          cUpdate.FileRoute,
		"clientProviderName": cUpdate.ClientProviderName,
	}}

	upd, err := db.OffersRequestsCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// FindByCount => devuelve las Ofertas filtrados  ---------- TotalOfferByCodeCompanyQuery() --- concurrente
func FindByCountAndSort(codeCompany string, count int, order string, typ string, page int) ([]*models.OfferRequest, int64, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	c := make(chan int64)

	var sort int = 1
	if typ == "desc" {
		sort = -1
	}
	pageNumber := 0
	if page > 1 {
		pageNumber = (page - 1) * count
	}

	condition := bson.M{"codeCompany": codeCompany}

	var offers []*models.OfferRequest
	cursor, err := db.OffersRequestsCollection.Find(ctx, condition, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return offers, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return offers, 0, false
	}

	go TotalOfferByCodeCompanyQuery(c, codeCompany)

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var offer models.OfferRequest
		err := cursor.Decode(&offer)
		if err != nil {
			return offers, 0, false
		}
		offer.ClientProviderName, _ = clientproviderservice.FindNameByCustID(offer.CodeReeup)
		offers = append(offers, &offer)
	}
	return offers, <-c, true
}
func TotalOfferByCodeCompanyQuery(c chan int64, code string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": code}

	cursor, err := db.OffersRequestsCollection.CountDocuments(ctx, condition)
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

func TotalOffersByCodeCompanyQueryClasif(code string) (int64, int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition := bson.M{"codeCompany": code, "state": "Activo"}
	condition2 := bson.M{"codeCompany": code, "state": "Inactivo"}

	active, err := db.OffersRequestsCollection.CountDocuments(ctx, condition)
	if err != nil {
		return 0, 0
	}
	inactive, err2 := db.OffersRequestsCollection.CountDocuments(ctx, condition2)
	if err2 != nil {
		return 0, 0
	}

	return active, inactive

}
