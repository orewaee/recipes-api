package services

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"github.com/rs/zerolog"
	"time"
)

type GuideService struct {
	guideRepo repos.GuideRepo
	cacheRepo repos.CacheRepo
	logger    *zerolog.Logger
}

func NewGuideService(repo repos.GuideRepo, cache repos.CacheRepo, logger *zerolog.Logger) apis.GuideApi {
	return &GuideService{repo, cache, logger}
}

func (service *GuideService) AddGuide(ctx context.Context, id, markdown string) error {
	service.logger.Log().Str("id", id).Msg("new guide")

	if err := service.guideRepo.AddGuide(ctx, id, markdown); err != nil {
		service.logger.Error().Err(err).Send()
		return err
	}

	return nil
}

func (service *GuideService) GetGuideById(ctx context.Context, id string) (string, error) {
	key := "guide_" + id
	if cache, err := service.cacheRepo.Get(ctx, key); err == nil {
		return cache, err
	}

	guide, err := service.guideRepo.GetGuideById(ctx, id)
	if err != nil {
		service.logger.Error().Err(err).Send()
		return "", err
	}

	service.cacheRepo.Put(ctx, key, guide, time.Minute*5)
	return guide, nil
}
