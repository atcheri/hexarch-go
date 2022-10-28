package adapters

import (
	"context"

	"github.com/atcheri/hexarch-go/internal/core/domain"
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
)

type inMemoryTranslations struct {
	db *databases.InMemoryDB
}

// NewInMemoryTranslations instantiates a new inMemorySentences that implements TranslationsRepository interface
func NewInMemoryTranslations(db *databases.InMemoryDB) ports.TranslationsRepository {
	return inMemoryTranslations{db: db}
}

func (i inMemoryTranslations) GetForProject(ctx context.Context, name string, offset, limit int) ([]domain.Translation, error) {
	return i.db.GetProjectTranslations(ctx, name, offset, limit)
}

func (i inMemoryTranslations) AddForProject(ctx context.Context, name, key, code, text string) error {
	return i.db.AddProjectTranslation(ctx, name, key, code, text)
}

func (i inMemoryTranslations) EditForProject(ctx context.Context, id, name, key, code, text string) error {
	return i.db.EditProjectTranslation(ctx, id, name, key, code, text)
}

func (i inMemoryTranslations) DeleteKeyForProject(_ context.Context, _, _ string) error {
	//TODO implement me
	panic("implement me")
}

func (i inMemoryTranslations) DeleteLanguageTranslationForProject(_ context.Context, _, _, _ string) error {
	//TODO implement me
	panic("implement me")
}
