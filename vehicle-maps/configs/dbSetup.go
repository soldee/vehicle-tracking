package configs

import (
    "context"
    "log"
    "time"
	"os"
	"github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {

	godotenv.Load(".env")
	uri := os.Getenv("MONGODB_URI");
	if (uri == "") {
		log.Fatal("MONGODB_URI env variable not found")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

var mongoClient = DBInstance()

func GetCollection(collectionName string) *mongo.Collection {
	return mongoClient.Database("VEHICLE-TRACKING").Collection(collectionName)
}
