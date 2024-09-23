package services

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/repos"
)

type GuideService struct {
	repo repos.GuideRepo
}

func NewGuideService(repo repos.GuideRepo) apis.GuideApi {
	return &GuideService{repo}
}

func (service *GuideService) AddGuide(ctx context.Context, id, markdown string) error {
	return service.repo.AddGuide(ctx, id, markdown)
}

func (service *GuideService) GetGuideById(ctx context.Context, id string) (string, error) {
	return service.repo.GetGuideById(ctx, id)
}
