package minio

import (
	"github.com/minio/minio-go"
	"io"
)

func (c *Connection) PutObject(filename string, data *io.Reader) error {
	contentType := "text/plain"

	_, err := c.client.PutObject("log-files", filename, *data, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	return nil
}

func (c *Connection) GetObject(filename string) ([]byte, error) {
	object, err := c.client.GetObject("log-files", filename, minio.GetObjectOptions{})
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
	err := c.client.RemoveObject("log-files", filename)
	if err != nil {
		return err
	}

	return nil
}
