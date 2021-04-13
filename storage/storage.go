package storage

import (
	"context"
	"io"
	"io/ioutil"

	"cloud.google.com/go/storage"
)

func SaveFile(file io.Reader, path string, bucket *storage.BucketHandle) (err error) {
	w := bucket.Object(path).NewWriter(context.Background())
	if _, err = io.Copy(w, file); err != nil {
		return err
	}
	return w.Close()
}

func GetFile(filePath string, bucket *storage.BucketHandle) (byteFile []byte, err error) {
	r, err := bucket.Object(filePath).NewReader(context.Background())
	if err != nil {
		return byteFile, err
	}
	defer r.Close()
	return ioutil.ReadAll(r)
}
