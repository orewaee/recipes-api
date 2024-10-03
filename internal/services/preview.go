package services

import (
	"context"
	"github.com/orewaee/recipes-api/internal/app/apis"
	"github.com/orewaee/recipes-api/internal/app/repos"
	"github.com/rs/zerolog"
)

type PreviewService struct {
	previewRepo repos.PreviewRepo
	logger      *zerolog.Logger
}

func NewPreviewService(previewRepo repos.PreviewRepo, logger *zerolog.Logger) apis.PreviewApi {
	return &PreviewService{previewRepo, logger}
}

func (service *PreviewService) AddPreview(ctx context.Context, id string, preview []byte) error {
	service.logger.Log().Str("id", id).Msg("new preview")

	if err := service.previewRepo.AddPreview(ctx, id, preview); err != nil {
		service.logger.Error().Err(err).Send()
		return err
	}

	return nil
}

func (service *PreviewService) GetPreviewById(ctx context.Context, id string) ([]byte, error) {
	return service.previewRepo.GetPreviewById(ctx, id)
}
