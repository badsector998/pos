package domain

import "github.com/gofrs/uuid"

type Item struct {
	BaseModel
	Name       string
	Price      float64
	Qty        int
	MerchantID uuid.NullUUID
	Merchant   Merchant
}

type ItemFilter struct {
	Name string
}

type ItemFilterOpts func(f *ItemFilter)

func (o *ItemFilterOpts) WithName(name string) ItemFilterOpts {
	return func(f *ItemFilter) {
		if name != "" {
			f.Name = name
		}
	}
}
