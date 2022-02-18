package categoryservice

import (
	"WEBCONTRACT-api-mongodb/db"
	"WEBCONTRACT-api-mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// FindByCount => devuelve las Categorias filtradas  ---------- TotalCategoriesQuery() --- concurrentes
func FindByCountAndSort(count int, order string, typ string, page int) ([]*models.Category, int64, bool) {

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

	var categories []*models.Category
	cursor, err := db.CategoryCollection.Find(ctx, bson.M{}, options.Find().SetLimit(int64(count)), options.Find().SetSkip(int64(pageNumber)).SetSort(bson.M{order: sort}))
	if err != nil {
		return categories, 0, false
	}
	err = cursor.Err()
	if err != nil {
		return categories, 0, false
	}

	go TotalCategoriesQuery(c)

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var category models.Category
		err := cursor.Decode(&category)
		if err != nil {
			return categories, 0, false
		}
		categories = append(categories, &category)
	}
	return categories, <-c, true
}
func TotalCategoriesQuery(c chan int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := db.CategoryCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		c <- 0
	} else {
		c <- cursor
	}
}

// FindAllCategories => Devuelve todas las categorias
func FindAllCategories() ([]*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var categories []*models.Category
	cursor, err := db.CategoryCollection.Find(ctx, bson.M{})
	if err != nil {
		return categories, err
	}
	err = cursor.Err()
	if err != nil {
		return categories, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var category models.Category
		err := cursor.Decode(&category)
		if err != nil {
			return categories, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}

// FindByID => Funcion para obtener categoria por id
func FindByID(id string) (models.Category, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	categoryID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": categoryID}
	var result models.Category

	err := db.CategoryCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return result, false
	}
	return result, true
}

// ValidateIfExistByNameAndCode => validar si existe la categoria ---- NameQuery() y CodeQuery()--- concurrentes
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
	var result models.Category

	err := db.CategoryCollection.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		c <- false
	} else {
		c <- true
	}

}
func CodeQuery(code string, d chan bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	condition2 := bson.M{"codeCategory": code}
	var result models.Category

	err := db.CategoryCollection.FindOne(ctx, condition2).Decode(&result)
	if err != nil {
		d <- false
	} else {
		d <- true
	}
}

// InsertNewCategory => inserta nueva categoria en la base de datos
func InsertNewCategory(c models.Category) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := db.CategoryCollection.InsertOne(ctx, c)

	return err
}

// UpdateByID => Funcion para actualizar categoria por id
func UpdateByID(id string, categoryUpdate models.Category) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	categoryID, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": categoryID}

	update := bson.M{"$set": bson.M{
		"codeCategory": categoryUpdate.CodeCategory,
		"name":         categoryUpdate.Name,
	}}

	upd, err := db.CategoryCollection.UpdateOne(ctx, condition, update)
	return upd.ModifiedCount, err
}

// DeleteByID => Funcion para eliminar categoria por id
func DeleteByID(id string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	categoryID, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{"_id": categoryID}

	delRes, err := db.CategoryCollection.DeleteOne(ctx, condition)
	if err != nil {
		return false
	}
	return delRes.DeletedCount > 0

}
