package adapter

import (
	"context"

	"github.com/badsector998/pos/internal/domain"
	"github.com/badsector998/pos/internal/lib/transaction"
	"github.com/badsector998/pos/internal/repository"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type ItemAdapter struct {
	db *gorm.DB
}

func NewItemAdapter(db *gorm.DB) *ItemAdapter {
	return &ItemAdapter{db: db}
}

var _ repository.Items = &ItemAdapter{}

func (a *ItemAdapter) CreateItem(ctx context.Context, i domain.Item) error {
	client := transaction.TransactionFromCtx(ctx, a.db)

	return client.Save(&i).Error
}

func (a *ItemAdapter) UpdateItem(ctx context.Context, i domain.Item) error {
	client := transaction.TransactionFromCtx(ctx, a.db)

	return client.Save(&i).Error
}

func (a *ItemAdapter) DeleteItem(ctx context.Context, i domain.Item) error {
	client := transaction.TransactionFromCtx(ctx, a.db)

	return client.Delete(&i).Error
}

func (a *ItemAdapter) ReadItemByID(ctx context.Context, id uuid.NullUUID) (domain.Item, error) {
	return domain.Item{}, nil
}

func (a *ItemAdapter) ReadItemByFilter(ctx context.Context, filter domain.ItemFilter) ([]domain.Item, error) {
	return nil, nil
}
