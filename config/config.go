package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDBga ulanishda xatolik:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDBni 'ping' qilishda xatolik:", err)
	}

	DB = client
	fmt.Println("MongoDBga muvaffaqiyatli ulandik!")
}

func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("quiz").Collection(collectionName)
}
