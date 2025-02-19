package repository

import (
	"context"
	"database/sql"

	"github.com/luispfcanales/confericis-backend/model"
	"github.com/luispfcanales/confericis-backend/ports/output"
)

type userRepository struct {
	db *sql.DB
}

func (userrepository *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	panic("not implemented") // TODO: Implement
}
func (userrepository *userRepository) Update(ctx context.Context, user *model.User) error {
	panic("not implemented") // TODO: Implement
}
func (userrepository *userRepository) Delete(ctx context.Context, id int64) error {
	panic("not implemented") // TODO: Implement
}

func NewUserRepository(db *sql.DB) output.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `
        INSERT INTO users (email, password, name, role_id)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at, updated_at`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.Password,
		user.Name,
		user.RoleID,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (r *userRepository) GetByID(ctx context.Context, id string) (*model.User, error) {
	query := `
        SELECT u.id, u.email, u.password, u.name, u.role_id,
               u.created_at, u.updated_at,
               r.id, r.name, r.description
        FROM users u
        JOIN roles r ON u.role_id = r.id
        WHERE u.id = $1`

	var user model.User
	var role model.Role
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Password, &user.Name, &user.RoleID,
		&user.CreatedAt, &user.UpdatedAt,
		&role.ID, &role.Name, &role.Description,
	)
	if err != nil {
		return nil, err
	}

	user.Role = &role
	return &user, nil
}
