package ports

import (
	"context"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

// TranslationsRepository defines the translations' repository interface
type TranslationsRepository interface {
	GetForProject(ctx context.Context, name string, offset, limit int) ([]domain.Translation, error)
	GetOneForProject(ctx context.Context, id string) (domain.Translation, error)
	AddForProject(ctx context.Context, name, key, code, text string) error
	EditForProject(ctx context.Context, id, key, code, text string) error
	DeleteKeyForProject(ctx context.Context, name, key string) error
	DeleteLanguageTranslationForProject(ctx context.Context, name, key, code string) error
}
