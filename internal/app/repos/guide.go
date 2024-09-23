package repos

import "context"

type GuideRepo interface {
	AddGuide(ctx context.Context, id, markdown string) error
	GetGuideById(ctx context.Context, id string) (string, error)
}
