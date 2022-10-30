package ports

import (
	"context"
)

type ProjectsRepository interface {
	Create(ctx context.Context, name string) error
	Edit(ctx context.Context, oldName, newName string) error
}
