package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	"github.com/IanDex/twitter/models"
)

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

// EncryptPass s
func EncryptPass(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}

// IntentoLogin IntentoLogin
func IntentoLogin(email string, password string) (models.Users, bool) {
	user, encontrado, _ := CheckDuplicateUser(email)
	if !encontrado {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
