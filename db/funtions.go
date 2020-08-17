package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/IanDex/twitter/models"
)

// EncryptPass s
func EncryptPass(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}

// CheckDuplicateUser UseruplicateUser
func CheckDuplicateUser(email string) (models.Users, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	condicion := bson.M{"email": email}

	var results models.Users
	err := col.FindOne(ctx, condicion).Decode(&results)
	ID := results.ID.Hex()
	if err != nil {
		return results, false, ID
	}

	return results, true, ID
}
