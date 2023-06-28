package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Connection struct {
	database *mongo.Database
}

var db *Connection

func NewClient() (dbConn *Connection, err error) {
	mgoUrl := "mongodb://localhost:27017"
	dbName := "local"
	timeout := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mgoUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	db = &Connection{database: client.Database(dbName)}

	return db, nil
}

func GetClient() *Connection {
	if db == nil {
		db, err := NewClient()
		if err != nil {
			return nil
		}

		return db
	}

	return db
}
