package services

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"time"
)

type GuideService struct {
	repo  repos.GuideRepo
	cache repos.CacheRepo
}

func NewGuideService(repo repos.GuideRepo, cache repos.CacheRepo) apis.GuideApi {
	return &GuideService{repo, cache}
}

func (service *GuideService) AddGuide(ctx context.Context, id, markdown string) error {
	return service.repo.AddGuide(ctx, id, markdown)
}

func (service *GuideService) GetGuideById(ctx context.Context, id string) (string, error) {
	key := "guide_" + id
	if cache, err := service.cache.Get(ctx, key); err == nil {
		return cache, err
	}

	guide, err := service.repo.GetGuideById(ctx, id)
	if err != nil {
		return "", err
	}

	service.cache.Put(ctx, key, guide, time.Minute*10)
	return guide, nil
}
