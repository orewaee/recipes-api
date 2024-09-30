package services

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"time"
)

type GuideService struct {
	guideRepo repos.GuideRepo
	cacheRepo repos.CacheRepo
}

func NewGuideService(repo repos.GuideRepo, cache repos.CacheRepo) apis.GuideApi {
	return &GuideService{repo, cache}
}

func (service *GuideService) AddGuide(ctx context.Context, id, markdown string) error {
	return service.guideRepo.AddGuide(ctx, id, markdown)
}

func (service *GuideService) GetGuideById(ctx context.Context, id string) (string, error) {
	key := "guide_" + id
	if cache, err := service.cacheRepo.Get(ctx, key); err == nil {
		return cache, err
	}

	guide, err := service.guideRepo.GetGuideById(ctx, id)
	if err != nil {
		return "", err
	}

	service.cacheRepo.Put(ctx, key, guide, time.Minute*5)
	return guide, nil
}
