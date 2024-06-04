package domain

type Merchant struct {
	BaseModel
	Name        string
	Address     string
	Owner       string
	Description string
}
