package ports

import (
	"context"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

type TranslationsRepository interface {
	GetForProject(ctx context.Context, name string, offset, limit int) ([]domain.Translation, error)
}
