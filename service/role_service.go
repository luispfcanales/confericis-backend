package service

import (
	"context"
	"errors"

	"github.com/luispfcanales/confericis-backend/model"
	"github.com/luispfcanales/confericis-backend/ports/input"
	"github.com/luispfcanales/confericis-backend/ports/output"
)

var (
	ErrInvalidRole = errors.New("invalid role")
)

type roleCaseUse struct {
	roleRepo output.RoleRepository
}

func NewRoleCaseUse(rr output.RoleRepository) input.RoleService {
	return &roleCaseUse{
		roleRepo: rr,
	}
}

func (s *roleCaseUse) GetRoles(ctx context.Context) ([]*model.Roles, error) {
	roles, err := s.roleRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return roles, nil
}
