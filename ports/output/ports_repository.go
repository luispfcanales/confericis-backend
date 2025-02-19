package output

import (
	"context"

	"github.com/luispfcanales/confericis-backend/model"
)

type RoleRepository interface {
	Create(ctx context.Context, role *model.Role) error
	GetByID(ctx context.Context, id string) (*model.Role, error)
	List(ctx context.Context) ([]*model.Role, error)
}
