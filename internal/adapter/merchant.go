package adapter

import (
	"context"

	"github.com/badsector998/pos/internal/domain"
	"github.com/badsector998/pos/internal/lib/transaction"
	"github.com/badsector998/pos/internal/repository"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type MerchantAdapter struct {
	db *gorm.DB
}

func NewMerchantAdapter(db *gorm.DB) *MerchantAdapter {
	return &MerchantAdapter{db: db}
}

var _ repository.Merchant = &MerchantAdapter{}

func (a *MerchantAdapter) CreateMerchant(ctx context.Context, m domain.Merchant) error {
	client := transaction.TransactionFromCtx(ctx, a.db)
	return client.Save(&m).Error
}

func (a *MerchantAdapter) UpdateMerchant(ctx context.Context, m domain.Merchant) error {
	client := transaction.TransactionFromCtx(ctx, a.db)
	return client.Save(&m).Error
}

func (a *MerchantAdapter) DeleteMerchant(ctx context.Context, m domain.Merchant) error {
	client := transaction.TransactionFromCtx(ctx, a.db)
	return client.Delete(&m).Error
}

func (a *MerchantAdapter) ReadMerchantByID(ctx context.Context, id uuid.NullUUID) (domain.Merchant, error) {
	return domain.Merchant{}, nil
}

func (a *MerchantAdapter) ReadMerchantByFilter(ctx context.Context) ([]domain.Merchant, error) {
	return nil, nil
}
