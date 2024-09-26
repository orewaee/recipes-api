package repos

import "context"

type PreviewRepo interface {
	AddPreview(ctx context.Context, id string, preview []byte) error
	GetPreviewById(ctx context.Context, id string) ([]byte, error)
}
