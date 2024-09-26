package services

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/repos"
)

type PreviewService struct {
	previewRepo repos.PreviewRepo
}

func NewPreviewService(previewRepo repos.PreviewRepo) apis.PreviewApi {
	return &PreviewService{previewRepo}
}

func (service *PreviewService) AddPreview(ctx context.Context, id string, preview []byte) error {
	return service.previewRepo.AddPreview(ctx, id, preview)
}

func (service *PreviewService) GetPreviewById(ctx context.Context, id string) ([]byte, error) {
	return service.previewRepo.GetPreviewById(ctx, id)
}
