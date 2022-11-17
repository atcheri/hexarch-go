package ports

import (
	"context"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

type CommentsRepository interface {
	GetAll(ctx context.Context, key string) ([]domain.Comment, error)
	Add(ctx context.Context, key, content string) error
	Edit(ctx context.Context, key, content string) error
}
