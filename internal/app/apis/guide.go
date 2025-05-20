package apis

import "context"

type GuideApi interface {
	SetGuide(ctx context.Context, id, markdown string) error
	GetGuideById(ctx context.Context, id string) (string, error)
}
