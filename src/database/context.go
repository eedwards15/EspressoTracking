package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseContext struct {
	client  *mongo.Client   // embeds the mongo.Client
	context context.Context // embeds the context.Context
}

func (dc *DatabaseContext) NewConnection() (mongo.Client, error) {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	//port := os.Getenv("DB_PORT")
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s", user, pass, host)

	var err error
	dc.client, err = mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Fatal(err)
	}

	dc.context, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = dc.client.Connect(dc.context)
	if err != nil {
		log.Fatal(err)
	}

	return *dc.client, nil
}

func (dc *DatabaseContext) Disconnect() {
	dc.client.Disconnect(dc.context)
	dc.context.Done()
	context.Background().Done()
}
