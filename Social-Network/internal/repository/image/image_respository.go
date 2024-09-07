package image

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/minio/minio-go/v7"
)

type ImageStorageRepository struct {
	Client *minio.Client
	Bucket string
}

func CreateImageStorageRepository(client *minio.Client,
	bucket string) *ImageStorageRepository {
	return &ImageStorageRepository{
		Client: client,
		Bucket: bucket,
	}
}

func (s *ImageStorageRepository) BucketExists() error {
	ctx := context.Background()
	exists, errBucketExists := s.Client.BucketExists(ctx, s.Bucket)

	if errBucketExists != nil {
		return errBucketExists
	}

	if exists == false {
		err := errors.New("Bucket is not existed")
		return err
	}

	return nil
}

func (s *ImageStorageRepository) getKey(postId int, filename string) string {
	return "post/" + strconv.Itoa(postId) + "/" + filename
}

func (s *ImageStorageRepository) PutImage(reader *os.File, postId int, filename string, fileSize int64) (string, error) {
	objectKey := s.getKey(postId, filename)

	info, err := s.Client.PutObject(
		context.Background(), s.Bucket, objectKey,
		reader, fileSize, minio.PutObjectOptions{},
	)

	if err != nil {
		return "", err
	}

	return info.Key, nil
}

func (s *ImageStorageRepository) GetSignedUrl(path string, expiration time.Duration) (string, error) {
	if path == "" {
		return "", nil
	}
	reqParams := make(url.Values)

	presignedURL, err := s.Client.PresignedGetObject(
		context.Background(), s.Bucket, path,
		expiration, reqParams,
	)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return presignedURL.String(), nil
}
func (s *ImageStorageRepository) DeleteImage(path string) error {
	// TODO
	return nil
}
