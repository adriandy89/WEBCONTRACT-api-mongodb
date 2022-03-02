package supplementservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertNewSuplement
func InsertNewSuplement(u models.Supplement) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.SuplementCollection.InsertOne(ctx, u)

	return err
}

// ValidateIfExistByCodeSupplementAndCodeContract
func ValidateIfExistByCodeContractAndCodeSupplement(codeCompany string, codeContract string, codeReeup string, codeSupplement string) bool {

	supplements, err := FindAllByCodeCompanyContractReeup(codeCompany, codeContract, codeReeup)
	if !err {
		return false
	}

	for i := 0; i < len(supplements); i++ {
		if supplements[i].CodeSupplement == codeSupplement {
			return true
		}
	}

	return false
}

// FindByID
func FindByID(id string) (models.Supplement, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": ID}
	var result models.Supplement

	err := db.SuplementCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// FindAllByCodeContract => return all suplements registered by Company
func FindAllByCodeCompanyContractReeup(codeCompany string, codeContract string, codeReeup string) ([]*models.Supplement, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*models.Supplement

	condition := bson.M{"codeCompany": codeCompany, "codeContract": codeContract, "codeReeup": codeReeup}
	cursor, err := db.SuplementCollection.Find(ctx, condition, options.Find().SetSort(bson.M{"codeSupplement": 1}))
	if err != nil {
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var suplement models.Supplement
		err := cursor.Decode(&suplement)
		if err != nil {
			return results, false
		}
		results = append(results, &suplement)
	}

	return results, true
}

// DeleteByID => Funcion para eliminar usuario por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	suplementID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": suplementID}

	delRes, err := db.SuplementCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}

// UpdateByID
func UpdateByID(id string, cUpdate models.Supplement) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": ID}

	update := bson.M{"$set": bson.M{
		"codeContract":   cUpdate.CodeContract,
		"codeReeup":      cUpdate.CodeReeup,
		"codeSupplement": cUpdate.CodeSupplement,
		"createdAt":      cUpdate.CreatedAt,
		"expireAt":       cUpdate.ExpireAt,
		"state":          cUpdate.State,
		"description":    cUpdate.Description,
		"ammountMN":      cUpdate.AmmountMN,
		"ammountCUC":     cUpdate.AmmountCUC,
		"operationMN":    cUpdate.OperationMN,
		"operationCUC":   cUpdate.OperationCUC,
		"codeCompany":    cUpdate.CodeCompany,
		"fileRoute":      cUpdate.FileRoute,
	}}

	upd, err := db.SuplementCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}
