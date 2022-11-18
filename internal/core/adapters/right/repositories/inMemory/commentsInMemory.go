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

func (i inMemoryComments) GetAll(ctx context.Context, id string) ([]domain.Comment, error) {
	return i.db.GetTranslationComments(ctx, id)
}

func (i inMemoryComments) Add(ctx context.Context, id, userID, text string) error {
	return i.db.AddProjectTranslationComment(ctx, id, userID, text)
}

func (i inMemoryComments) Edit(ctx context.Context, id, userID, text string) error {
	//TODO implement me
	panic("implement me")
}
