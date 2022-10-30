package adapters

import (
	"context"

	ports "github.com/atcheri/hexarch-go/internal/core/ports/right/repositories"
	"github.com/atcheri/hexarch-go/internal/infrastructure/databases"
)

type inMemoryProjects struct {
	db *databases.InMemoryDB
}

// NewInMemoryProjects instantiates a new inMemoryProjects that implements ProjectsRepository interface
func NewInMemoryProjects(db *databases.InMemoryDB) ports.ProjectsRepository {
	return inMemoryProjects{db: db}
}

func (i inMemoryProjects) Create(ctx context.Context, name string) error {
	return i.db.CreateProject(ctx, name)
}

func (i inMemoryProjects) Edit(ctx context.Context, oldName, newName string) error {
	return i.db.EditProject(ctx, oldName, newName)
}
