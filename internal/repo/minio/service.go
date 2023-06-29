package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
)

func (c *Connection) PutObject(filename string, data []byte) error {
	bytesReader := bytes.NewReader(data)
	reader := io.Reader(bytesReader)

	_, err := c.client.PutObject(context.Background(), "log-files", filename, reader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) GetObject(filename string) ([]byte, error) {
	object, err := c.client.GetObject(context.Background(), "log-files", filename, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer object.Close()

	data, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *Connection) DeleteObject(filename string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
		VersionID:        "",
	}
	err := c.client.RemoveObject(context.Background(), "log-files", filename, opts)
	if err != nil {
		return err
	}

	return nil
}
