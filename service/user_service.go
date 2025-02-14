package service

import (
	"context"
	"errors"

	"github.com/confericis-backend/model"
	"github.com/confericis-backend/ports/input"
	"github.com/confericis-backend/ports/output"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailExists = errors.New("email already exists")
	ErrInvalidRole = errors.New("invalid role")
)

type userService struct {
	userRepo output.UserRepository
	roleRepo output.RoleRepository
}

func (userservice *userService) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}
func (userservice *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}
func (userservice *userService) UpdateUser(ctx context.Context, user *model.User) error {
	panic("not implemented") // TODO: Implement
}
func (userservice *userService) DeleteUser(ctx context.Context, id int64) error {
	panic("not implemented") // TODO: Implement
}

func NewUserService(ur output.UserRepository, rr output.RoleRepository) input.UserService {
	return &userService{
		userRepo: ur,
		roleRepo: rr,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) error {
	// Validar que el email no exista
	existingUser, err := s.userRepo.GetByEmail(ctx, user.Email)
	if err == nil && existingUser != nil {
		return ErrEmailExists
	}

	// Validar que el rol exista
	role, err := s.roleRepo.GetByID(ctx, user.RoleID)
	if err != nil {
		return ErrInvalidRole
	}
	user.Role = role

	// Encriptar password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userRepo.Create(ctx, user)
}
