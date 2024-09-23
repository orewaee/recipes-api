package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"io"
	"strings"
)

type GuideRepo struct {
	bucket string
	client *minio.Client
}

func NewGuideRepo(ctx context.Context, endpoint, user, password, bucket string) (repos.GuideRepo, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

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

	return &GuideRepo{bucket, client}, nil
}

func (repo *GuideRepo) AddGuide(ctx context.Context, id, markdown string) error {
	reader := strings.NewReader(markdown)
	_, err := repo.client.PutObject(ctx, repo.bucket, id+".md", reader, reader.Size(), minio.PutObjectOptions{
		ContentType: "text/markdown; charset=UTF-8",
	})

	return err
}

func (repo *GuideRepo) GetGuideById(ctx context.Context, id string) (string, error) {
	object, err := repo.client.GetObject(ctx, repo.bucket, id+".md", minio.GetObjectOptions{})
	if err != nil {
		return "", err
	}

	bytes, err := io.ReadAll(object)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
