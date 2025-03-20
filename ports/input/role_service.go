package input

import (
	"context"

	"github.com/luispfcanales/confericis-backend/model"
)

type RoleService interface {
	// CreateUser(ctx context.Context, user *model.User) error
	GetRoleByID(ctx context.Context, id string) (*model.Roles, error)
	GetRoles(ctx context.Context) ([]*model.Roles, error)
	// UpdateUser(ctx context.Context, user *model.User) error
	// DeleteUser(ctx context.Context, id int64) error
}
