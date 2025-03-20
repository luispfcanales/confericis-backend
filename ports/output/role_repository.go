package output

import (
	"context"

	"github.com/luispfcanales/confericis-backend/model"
)

type RoleRepository interface {
	// Create(ctx context.Context, role *model.Roles) error
	GetByID(ctx context.Context, id string) (*model.Roles, error)
	List(ctx context.Context) ([]*model.Roles, error)
	Update(ctx context.Context, role *model.Roles) error
}
