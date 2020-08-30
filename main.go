package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	myDatabase := client.Database("test")
	myCollection := myDatabase.Collection("aliases")

	filterCursor, err := myCollection.Find(ctx, bson.M{"KEY": "SPR1011"})
	if err != nil {
		log.Fatal(err)
	}
	var aliasFiltered []bson.M
	if err = filterCursor.All(ctx, &aliasFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(aliasFiltered)
}
