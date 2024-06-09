package adapter

import (
	"context"

	"github.com/badsector998/pos/internal/domain"
	"github.com/badsector998/pos/internal/lib/transaction"
	"github.com/badsector998/pos/internal/repository"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserAdapter struct {
	db *gorm.DB
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{db: db}
}

var _ repository.User = &UserAdapter{}

func (a *UserAdapter) CreateUser(ctx context.Context, u domain.User) error {
	client := transaction.TransactionFromCtx(ctx, a.db)
	return client.Save(&u).Error
}

func (a *UserAdapter) UpdateUser(ctx context.Context, u domain.User) error {
	client := transaction.TransactionFromCtx(ctx, a.db)
	return client.Save(&u).Error
}

func (a *UserAdapter) DeleteUser(ctx context.Context, u domain.User) error {
	client := transaction.TransactionFromCtx(ctx, a.db)
	return client.Delete(&u).Error
}

func (a *UserAdapter) ReadUserByID(ctx context.Context, id uuid.NullUUID) (domain.User, error) {
	return domain.User{}, nil
}

func (a *UserAdapter) ReadUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return domain.User{}, nil
}

func (a *UserAdapter) ReadUserByFilter(ctx context.Context) ([]domain.User, error) {
	return nil, nil
}
