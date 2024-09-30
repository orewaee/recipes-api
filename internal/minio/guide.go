package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/orewaee/recipes-api/internal/app/domain"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"io"
	"strings"
)

type GuideRepo struct {
	bucket string
	client *minio.Client
}

func NewGuideRepo(ctx context.Context, client *minio.Client, bucket string) (repos.GuideRepo, error) {
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

	guide, err := io.ReadAll(object)
	if err != nil && err.Error() == "The specified key does not exist." {
		return "", domain.ErrNoGuide
	}

	if err != nil {
		return "", err
	}

	return string(guide), nil
}
