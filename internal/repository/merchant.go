package repository

import (
	"context"

	"github.com/badsector998/pos/internal/domain"
	"github.com/gofrs/uuid"
)

type Merchant interface {
	CreateMerchant(ctx context.Context, merchant domain.Merchant) error
	UpdateMerchant(ctx context.Context, merchant domain.Merchant) error
	DeleteMerchant(ctx context.Context, merchant domain.Merchant) error
	ReadMerchantByID(ctx context.Context, id uuid.NullUUID) (domain.Merchant, error)
	ReadMerchantByFilter(ctx context.Context)
}
