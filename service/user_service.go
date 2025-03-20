package service

import (
	"context"
	"errors"

	"github.com/luispfcanales/confericis-backend/model"
	"github.com/luispfcanales/confericis-backend/ports/input"
	"github.com/luispfcanales/confericis-backend/ports/output"
)

var (
	ErrEmailExists = errors.New("email already exists")
)

type userCaseUse struct {
	userRepo output.UserRepository
	roleRepo output.RoleRepository
}

func (userservice *userCaseUse) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}
func (userservice *userCaseUse) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}
func (userservice *userCaseUse) UpdateUser(ctx context.Context, user *model.User) error {
	panic("not implemented") // TODO: Implement
}
func (userservice *userCaseUse) DeleteUser(ctx context.Context, id int64) error {
	panic("not implemented") // TODO: Implement
}

func NewUserCaseUse(ur output.UserRepository, rr output.RoleRepository) input.UserService {
	return &userCaseUse{
		userRepo: ur,
		roleRepo: rr,
	}
}

func (s *userCaseUse) CreateUser(ctx context.Context, user *model.User) error {
	// Validar que el email no exista
	return nil
	// existingUser, err := s.userRepo.GetByEmail(ctx, user.Email)
	// if err == nil && existingUser != nil {
	// 	return ErrEmailExists
	// }

	// // Validar que el rol exista
	// role, err := s.roleRepo.GetByID(ctx, user.RoleID)
	// if err != nil {
	// 	return ErrInvalidRole
	// }
	// user.Role = role

	// // Encriptar password
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return err
	// }
	// user.Password = string(hashedPassword)

	// return s.userRepo.Create(ctx, user)
}
