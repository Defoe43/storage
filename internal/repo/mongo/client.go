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

// GetMongoConnection returns a handle to the given database
func GetMongoConnection(url, dbname string, timeout time.Duration) (dbConn *Connection, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if db == nil {
		clientOptions := options.Client().ApplyURI(url)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			return nil, err
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			return nil, err
		}

		db = &Connection{database: client.Database(dbname)}

		return db, nil
	} else if err := db.database.Client().Ping(ctx, nil); err != nil {
		db = nil
		db, err := GetMongoConnection(url, dbname, timeout)
		if err != nil {
			return nil, err
		}

		return db, nil
	}

	return db, nil
}

func GetDatabase() *Connection {
	return db
}
