package database

import (
	"context"
	"log"
	"time"

	"golang_ot/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient   *mongo.Client
	MongoDatabase *mongo.Database // ✅ Store database globally
)

func ConnectMongo(cfg config.MongoConfig) *mongo.Database {
	// Create a context with timeout for connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create MongoDB client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URL))
	if err != nil {
		log.Fatalf("❌ Failed to create MongoDB client: %v", err)
	}

	// Ping to ensure connection is working
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("❌ Failed to ping MongoDB: %v", err)
	}

	// Set global MongoClient
	MongoClient = client
	MongoDatabase = client.Database(cfg.Database)
	log.Println("✅ Connected to MongoDB successfully:", cfg.Database)

	// Return the database instance
	return client.Database(cfg.Database)
}

// GetCollection returns a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if MongoDatabase == nil {
		log.Fatal("❌ MongoDatabase is not initialized. Call ConnectMongo() first.")
	}
	return MongoDatabase.Collection(collectionName)
}
