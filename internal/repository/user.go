package repository

import (
	"context"

	"github.com/badsector998/pos/internal/domain"
	"github.com/gofrs/uuid"
)

type User interface {
	CreateUser(ctx context.Context, user domain.User) error
	UpdateUser(ctx context.Context, user domain.User) error
	DeleteUser(ctx context.Context, user domain.User) error
	ReadUserByID(ctx context.Context, id uuid.NullUUID) (domain.User, error)
	ReadUserByEmail(ctx context.Context, email string) (domain.User, error)
	ReadUserByFilter(ctx context.Context)
}
