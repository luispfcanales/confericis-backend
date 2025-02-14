package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/confericis-backend/model"
	"github.com/confericis-backend/ports/output"
)

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) output.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) Create(ctx context.Context, role *model.Role) error {
	query := `
        INSERT INTO roles (name, description, created_at, updated_at)
        VALUES ($1, $2, $3, $4)
        RETURNING id`

	now := time.Now()

	err := r.db.QueryRowContext(
		ctx,
		query,
		role.Name,
		role.Description,
		now,
		now,
	).Scan(&role.ID)

	if err != nil {
		return err
	}

	role.CreatedAt = now
	role.UpdatedAt = now
	return nil
}

func (r *roleRepository) GetByID(ctx context.Context, id string) (*model.Role, error) {
	query := `
        SELECT id, name, description, created_at, updated_at
        FROM roles
        WHERE id = $1`

	role := &model.Role{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("role not found")
	}

	if err != nil {
		return nil, err
	}

	return role, nil
}

func (r *roleRepository) List(ctx context.Context) ([]*model.Role, error) {
	query := `
        SELECT id, name, description, created_at, updated_at
        FROM roles
        ORDER BY name`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*model.Role
	for rows.Next() {
		role := &model.Role{}
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) Update(ctx context.Context, role *model.Role) error {
	query := `
        UPDATE roles 
        SET name = $1, 
            description = $2,
            updated_at = $3
        WHERE id = $4
        RETURNING created_at`

	now := time.Now()

	err := r.db.QueryRowContext(
		ctx,
		query,
		role.Name,
		role.Description,
		now,
		role.ID,
	).Scan(&role.CreatedAt)

	if err == sql.ErrNoRows {
		return errors.New("role not found")
	}

	if err != nil {
		return err
	}

	role.UpdatedAt = now
	return nil
}

func (r *roleRepository) Delete(ctx context.Context, id int64) error {
	query := `
        DELETE FROM roles 
        WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("role not found")
	}

	return nil
}
