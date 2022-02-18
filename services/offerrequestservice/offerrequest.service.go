package offerrequestservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertNewOfferRequest
func InsertNewOfferRequest(u models.OfferRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.OffersRequestsCollection.InsertOne(ctx, u)

	return err
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

// FindAllByCompany
func FindAllByCompany(company string) ([]*models.OfferRequest, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*models.OfferRequest

	condition := bson.M{"codeCompany": company}
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
		"codeReeup":   cUpdate.CodeReeup,
		"codeOffer":   cUpdate.CodeOffer,
		"ammountMN":   cUpdate.AmmountMN,
		"ammountCUC":  cUpdate.AmmountCUC,
		"state":       cUpdate.State,
		"description": cUpdate.Description,
		"codeCompany": cUpdate.CodeCompany,
		"fileRoute":   cUpdate.FileRoute,
	}}

	upd, err := db.OffersRequestsCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}
