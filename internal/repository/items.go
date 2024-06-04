package repository

import (
	"context"

	"github.com/badsector998/pos/internal/domain"
	"github.com/gofrs/uuid"
)

type Items interface {
	CreateItem(ctx context.Context, item domain.Item) error
	UpdateItem(ctx context.Context, item domain.Item) error
	DeleteItem(ctx context.Context, item domain.Item) error
	ReadItemByID(ctx context.Context, id uuid.NullUUID) (domain.Item, error)
	ReadItemByFilter(ctx context.Context, filter domain.ItemFilter) ([]domain.Item, error)
}
