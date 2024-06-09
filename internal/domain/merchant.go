package domain

import (
	"fmt"

	"gorm.io/gorm"
)

type Merchant struct {
	BaseModel
	Name        string
	Address     string
	Owner       string
	Description string
}

func (m *Merchant) BeforeUpdate(_ *gorm.DB) error {
	if !m.ID.Valid || m.ID.UUID.String() == "" {
		return fmt.Errorf("id is invalid")
	}

	return nil
}

func (m *Merchant) BeforeDelete(_ *gorm.DB) error {
	if !m.ID.Valid || m.ID.UUID.String() == "" {
		return fmt.Errorf("id is invalid")
	}

	return nil
}

type MerchantFilter struct {
	Name  string
	Owner string
}

type MerchantFilterOpts func(m *MerchantFilter)

func (o *MerchantFilterOpts) WithName(name string) MerchantFilterOpts {
	return func(m *MerchantFilter) {
		if name != "" {
			m.Name = name
		}
	}
}

func (o *MerchantFilterOpts) WithOwner(owner string) MerchantFilterOpts {
	return func(m *MerchantFilter) {
		if owner != "" {
			m.Owner = owner
		}
	}
}
