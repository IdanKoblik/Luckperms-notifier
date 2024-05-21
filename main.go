package main

import (
	"context"
	"log"
	"luckperms-notifier/config"
	"luckperms-notifier/mongo_watcher"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoURL, err := config.GetMongoURL(); if err != nil {
		log.Fatalf("%v\n", err)
	}
	
	clientOptions := options.Client().ApplyURI(mongoURL)
    client, err := mongo.Connect(context.Background(), clientOptions); if err != nil {
        log.Fatalf("Error connecting to MongoDB: %v\n", err)
        return
    }

	err = client.Ping(context.Background(), nil); if err != nil {
        log.Fatalf("Error pinging MongoDB: %v\n", err)
        return
    }

	databaseName, err := config.GetMongoDatabase(); if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

	mongoCollection, err := config.GetMongoCollection(); if err != nil {
		log.Fatalf("%v\n", err)
		return
	}

	log.Printf("Started system for db: %s\n", databaseName)
    collection := client.Database(databaseName).Collection(mongoCollection)
	options := options.ChangeStream().SetFullDocument(options.UpdateLookup)
	mongo_watcher.WatchCollection(collection, options)
}