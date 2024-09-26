package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"io"
)

type PreviewRepo struct {
	bucket string
	client *minio.Client
}

func NewPreviewRepo(ctx context.Context, client *minio.Client, bucket string) (repos.PreviewRepo, error) {
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	}

	if !exists {
		err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			return nil, err
		}
	}

	return &PreviewRepo{bucket, client}, nil
}

func (repo *PreviewRepo) AddPreview(ctx context.Context, id string, preview []byte) error {
	reader := bytes.NewReader(preview)
	_, err := repo.client.PutObject(ctx, repo.bucket, id, reader, reader.Size(), minio.PutObjectOptions{
		ContentType: "image/png",
	})

	return err
}

func (repo *PreviewRepo) GetPreviewById(ctx context.Context, id string) ([]byte, error) {
	object, err := repo.client.GetObject(ctx, repo.bucket, id, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return io.ReadAll(object)
}
