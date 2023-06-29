package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
)

func (c *Connection) GetGridFSBucket() (*gridfs.Bucket, error) {
	bucket, err := gridfs.NewBucket(
		c.database,
		options.GridFSBucket().SetName("log_files"),
	)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (c *Connection) PutFile(filename string, data []byte) error {
	bucket, err := c.GetGridFSBucket()
	if err != nil {
		log.Println(err)
	}

	uploadStream, err := bucket.OpenUploadStream(filename)
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

func (c *Connection) GetFile(filename string) ([]byte, error) {
	bucket, err := c.GetGridFSBucket()
	if err != nil {
		return nil, err
	}

	stream, err := bucket.OpenDownloadStreamByName(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer stream.Close()

	data, err := io.ReadAll(stream)
	if err != nil {
		log.Fatal(err)
	}

	return data, nil
}

func (c *Connection) DeleteFile(filename string) error {
	fsFiles := c.database.Collection("log_files.files")
	fsChunks := c.database.Collection("log_files.chunks")

	filter := bson.M{"filename": filename}
	var result bson.M
	err := fsFiles.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return err
	}

	fileId, ok := result["_id"].(primitive.ObjectID)
	if !ok {
		return fmt.Errorf("%s", err)
	}

	_, err = fsFiles.DeleteOne(context.Background(), bson.M{"_id": fileId})
	if err != nil {
		return err
	}

	_, err = fsChunks.DeleteMany(context.Background(), bson.M{"files_id": fileId})
	if err != nil {
		return err
	}

	return nil
}
