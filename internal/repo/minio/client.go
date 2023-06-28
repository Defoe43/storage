package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Connection struct {
	client *minio.Client
}

var objectStorage *Connection

func NewClient() (*Connection, error) {
	endpoint := "192.168.0.45:30960"
	//endpoint := "localhost:9006"
	accessKeyID := "3VZTiEVgzIf4oXn6RIym"
	//accessKeyID := "qwep12345"
	secretAccessKey := "wysjop-4nakQo-pycdav"
	//secretAccessKey := "qwep12345"

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		return nil, err
	}

	return &Connection{
		client: minioClient,
	}, nil
}

func GetClient() (*Connection, error) {
	if objectStorage == nil {
		objectStorage, err := NewClient()
		if err != nil {
			return nil, err
		}

		return objectStorage, nil
	}

	return objectStorage, nil
}
