package ports

import (
	"context"
)

type ProjectsRepository interface {
	Create(ctx context.Context, name string) error
}
