package repos

import "context"

type GuideRepo interface {
	SetGuide(ctx context.Context, id, markdown string) error
	GetGuideById(ctx context.Context, id string) (string, error)
}
