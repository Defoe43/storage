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

type Repository struct {
	database *mongo.Database
}

var db *Connection

func NewClient() (dbConn *Connection, err error) {
	mgoUrl := "mongodb://root:8Slaw_FluKnoc@192.168.0.45:30967/"
	//mgoUrl := "localhost:27017"
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

func newConnection() (*mongo.Database, error) {
	mgoUrl := "mongodb://root:8Slaw_FluKnoc@192.168.0.45:30967/"
	//mgoUrl := "localhost:27017"
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

	dbConn := client.Database(dbName)

	return dbConn, nil
}

func (r *Repository) Connect() error {
	db, err := newConnection()
	if err != nil {
		return err
	}

	r.database = db

	return nil
}

func (r *Repository) Close() error {
	err := r.database.Client().Disconnect(context.Background())
	return err
}

func (c *Connection) CloseConnection() error {
	err := c.database.Client().Disconnect(context.Background())
	return err
}

func GetClient() (db *Connection, err error) {
	if db == nil {
		db, err = NewClient()
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
