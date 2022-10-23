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
