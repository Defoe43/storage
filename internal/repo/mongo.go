package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (c *Connection) GetGridFS() (*gridfs.Bucket, error) {
	gfs, err := gridfs.NewBucket(
		c.database,
		options.GridFSBucket().SetName("log_files"),
	)
	if err != nil {
		return nil, err
	}

	return gfs, nil
}

func (c *Connection) PutFile(filename string, data []byte) error {
	gfs, err := c.GetGridFS()
	if err != nil {
		log.Println(err)
	}

	uploadStream, err := gfs.OpenUploadStream(filename)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) GetFile(filename string) (*gridfs.DownloadStream, error) {
	gfs, err := c.GetGridFS()
	if err != nil {
		return nil, err
	}

	stream, err := gfs.OpenDownloadStreamByName(filename)
	if err != nil {
		log.Fatal(err)
	}

	//err = stream.GetFile().UnmarshalBSON(file)
	//if err != nil {
	//	log.Println(err)
	//}

	//filter := bson.D{{"filename", filename}}
	//var file bson.M
	//err = gfs.GetFilesCollection().FindOne(context.Background(), filter).Decode(&file)
	//if err != nil {
	//	return nil, err
	//}
	//fileID := file["_id"].(primitive.ObjectID)
	//var b bytes.Buffer
	//_, err = gfs.DownloadToStream(fileID, &b)
	//if err != nil {
	//	return nil, err
	//}
	//
	//data, err := io.ReadAll(&b)
	//if err != nil {
	//	return nil, err
	//}

	return stream, nil
}
