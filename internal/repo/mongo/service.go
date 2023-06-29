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

func (r *Repository) getGridFSBucket() (*gridfs.Bucket, error) {
	bucket, err := gridfs.NewBucket(
		r.database,
		options.GridFSBucket().SetName("log_files"),
	)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}

func (r *Repository) PutFile(filename string, data []byte) error {
	bucket, err := r.getGridFSBucket()
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

func (r *Repository) GetFile(filename string) ([]byte, error) {
	bucket, err := r.getGridFSBucket()
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

func (r *Repository) DeleteFile(filename string) error {
	fsFiles := r.database.Collection("log_files.files")
	fsChunks := r.database.Collection("log_files.chunks")

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
