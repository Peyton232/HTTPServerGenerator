package database

import (
	"context"

	"github.com/Peyton232/HTTPServerGenerator/exampleServer/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) UpdatePet(newPet models.Pet) (models.Pet, error) {
	result := db.pet.FindOneAndUpdate(context.Background(), bson.M{"id": newPet.Id}, newPet)
	pet := models.Pet{}
	result.Decode(pet)
	return pet, result.Err()
}
