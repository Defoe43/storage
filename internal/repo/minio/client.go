package minio

import (
	"github.com/minio/minio-go"
)

type Connection struct {
	client *minio.Client
}

var objectStorage *Connection

func NewClient() (*Connection, error) {
	minioClient, err := minio.New("localhost:9006", "qwep12345", "qwep12345", false)
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
