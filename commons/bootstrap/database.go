package bootstrap

import (
	"context"
	"log"
	"time"
	"todo-list-api/infrastructures/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoURI := env.MongoURI
	client, err := mongo.NewClient(mongoURI)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func CloseMongoConnnection(client mongo.Client) {
	if client == nil {
		return
	}
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MongoDB connection closed")
}
