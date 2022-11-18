package adapters

import (
	"context"

	"github.com/atcheri/hexarch-go/internal/core/domain"
	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
)

type inMemoryComments struct {
	db *databases.InMemoryDB
}

// NewInMemoryComments instantiates a new inMemoryComments that implements CommentsRepository interface
func NewInMemoryComments(db *databases.InMemoryDB) ports.CommentsRepository {
	return inMemoryComments{db: db}
}

func (i inMemoryComments) GetAll(ctx context.Context, key string) ([]domain.Comment, error) {
	return i.db.GetTranslationComments(ctx, key)
}

func (i inMemoryComments) Add(ctx context.Context, key, content string) error {
	//TODO implement me
	panic("implement me")
}

func (i inMemoryComments) Edit(ctx context.Context, key, content string) error {
	//TODO implement me
	panic("implement me")
}
