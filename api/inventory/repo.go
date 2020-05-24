package inventory

import (
	"context"
	// "database/sql"
	"fmt"
	"log"

	// "strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repo ...
type Repo struct {
}

var connStr = "mongodb://localhost:27017"

//ProvideInventoryRepo ... provide for the DI
func ProvideInventoryRepo() Repo {
	return Repo{}
}

//GetClient ...
func (r *Repo) GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Inventory!")
	return client
}

//Add ... Adds an item
func (r *Repo) Add(item Item) string {
	clientOptions := options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}

	collection := client.Database("inventory").Collection("items")
	item.ID = primitive.NewObjectID()
	result, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		log.Println(err)
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}

//Find ... Gets an item
func (r *Repo) Find(ID string) []*Item {
	clientOptions := options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println(err)
	}

	id, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.D{{"_id", id}}
	var items []*Item

	collection := client.Database("inventory").Collection("items")
	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Println(err)
	}

	for cur.Next(context.TODO()) {
		var item Item
		err := cur.Decode(&item)
		if err != nil {
			log.Println(err)
		}

		items = append(items, &item)

	}

	return items
}

//Ping ...  Ping the DB to verify if the able to connect to the db
func (r *Repo) Ping() {
	clientOptions := options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Println(err)
	}
	fmt.Println("Connection to mongo server is all good!")
}
