package database

import (
	"context"

	"github.com/Peyton232/HTTPServerGenerator-exampleServer/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) DeletePetByID(id string) (models.Pet, error) {
	result := db.pet.FindOneAndDelete(context.Background(), bson.M{"id": id})
	pet := models.Pet{}
	result.Decode(pet)
	return pet, result.Err()
}

func (db *DB) DeletePetByName(name string) (models.Pet, error) {
	result := db.pet.FindOne(context.Background(), bson.M{"name": name})
	pet := models.Pet{}
	result.Decode(pet)
	return pet, result.Err()
}
