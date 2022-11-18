package ports

import (
	"context"

	"github.com/atcheri/hexarch-go/internal/core/domain"
)

// CommentsRepository defines the Comments' repository interface
type CommentsRepository interface {
	GetAll(ctx context.Context, id string) ([]domain.Comment, error)
	Add(ctx context.Context, id, userID, text string) error
	Edit(ctx context.Context, id, userID, text string) error
}
