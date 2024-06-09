package domain

import (
	"fmt"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	StaffRole   UserRole = "staff"
	CashierRole UserRole = "cashier"
)

type User struct {
	BaseModel
	Name       string
	Email      string
	Password   string
	Role       UserRole
	MerchantID uuid.NullUUID
	Merchant   Merchant
}

func (u *User) BeforeUpdate(_ *gorm.DB) error {
	if !u.ID.Valid || u.ID.UUID.String() == "" {
		return fmt.Errorf("id is invalid")
	}

	return nil
}

func (u *User) BeforeDelete(_ *gorm.DB) error {
	if !u.ID.Valid || u.ID.UUID.String() == "" {
		return fmt.Errorf("id is invalid")
	}

	return nil
}
