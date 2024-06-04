package domain

import "github.com/gofrs/uuid"

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
	Role       string
	MerchantID uuid.NullUUID
	Merchant   Merchant
}
