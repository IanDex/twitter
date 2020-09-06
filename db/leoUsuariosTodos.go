package db

import (
	"context"
	"github.com/IanDex/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func LeoUsuarioTodos(ID string, page int64, search string, tipo string) ([]*models.Users, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("users")

	var results []*models.Users

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var user models.Users
		err := cursor.Decode(&user)
		if err != nil {
			return results, false
		}

		var relacion models.Relacion
		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = user.ID.Hex()

		incluir = false

		encontrado, err = ConsultaRelacion(relacion)
		if tipo == "new" && !encontrado {
			incluir = true
		} else if tipo == "follow" && encontrado {
			incluir = true
		}

		if relacion.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			user.Password = ""
			user.Biografia = ""
			user.Biografia = ""
			user.SitioWeb = ""
			user.Ubicacion = ""
			user.Banner = ""
			user.Email = ""

			results = append(results, &user)
		}
	}

	if cursor.Err() != nil {
		return results, false
	}

	_ = cursor.Close(ctx)
	return results, true

}
