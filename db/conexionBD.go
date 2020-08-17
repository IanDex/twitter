package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN d*/
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:root@cluster0.qzlao.mongodb.net/twitter?retryWrites=true&w=majority")

/*ConectarDB d */
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión Exitosa")
	return client
}

/*CheckConection dsds */
func CheckConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
