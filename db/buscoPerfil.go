package db

import (
	"context"
	"fmt"
	"github.com/IanDex/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func BuscoPerfil(ID string) (models.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var perfil models.Users
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id" : objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""
	if err != nil {
		fmt.Printf("User not found " + err.Error())
		return perfil, err
	}

	return perfil, nil
}
