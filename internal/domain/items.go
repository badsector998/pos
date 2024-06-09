package domain

import (
	"fmt"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type Item struct {
	BaseModel
	Name       string
	Price      float64
	Qty        int
	MerchantID uuid.NullUUID
	Merchant   Merchant
}

func (i *Item) BeforeUpdate(_ *gorm.DB) error {
	if !i.ID.Valid || i.ID.UUID.String() == "" {
		return fmt.Errorf("id is invalid")
	}

	return nil
}

func (i *Item) BeforeDelete(_ *gorm.DB) error {
	if !i.ID.Valid || i.ID.UUID.String() == "" {
		return fmt.Errorf("id is invalid")
	}

	return nil
}

type ItemFilter struct {
	Name       string
	MerchantID uuid.NullUUID
}

type ItemFilterOpts func(f *ItemFilter)

func (o *ItemFilterOpts) WithName(name string) ItemFilterOpts {
	return func(f *ItemFilter) {
		if name != "" {
			f.Name = name
		}
	}
}

func (o *ItemFilterOpts) WithMerchantID(id uuid.NullUUID) ItemFilterOpts {
	return func(f *ItemFilter) {
		if id.Valid && id.UUID.String() != "" {
			f.MerchantID = id
		}
	}
}
