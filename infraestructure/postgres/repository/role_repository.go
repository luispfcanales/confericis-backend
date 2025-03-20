package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/luispfcanales/confericis-backend/model"
	"github.com/luispfcanales/confericis-backend/ports/output"
)

type roleRepository struct {
	db *sql.DB
}

func NewRoleRepository(db *sql.DB) output.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) Create(ctx context.Context, role *model.Roles) error {
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

func (r *roleRepository) GetByID(ctx context.Context, id string) (*model.Roles, error) {
	query := `
        SELECT id, name, status, created_at, updated_at, deleted_at
        FROM roles
        WHERE id = $1`

	role := &model.Roles{}
	var deletedAt sql.NullTime
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&role.ID,
		&role.Name,
		&role.Status,
		&role.CreatedAt,
		&role.UpdatedAt,
		&deletedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("role not found")
	}

	if err != nil {
		return nil, err
	}

	if deletedAt.Valid {
		role.DeletedAt = deletedAt.Time
	}

	return role, nil
}

func (r *roleRepository) List(ctx context.Context) ([]*model.Roles, error) {
	query := `
        SELECT id, name,status, created_at, updated_at, deleted_at
        FROM roles
        ORDER BY name`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*model.Roles
	for rows.Next() {
		role := &model.Roles{}
		var deletedAt sql.NullTime
		err := rows.Scan(
			&role.ID,
			&role.Name,
			&role.Status,
			// &role.Description,
			&role.CreatedAt,
			&role.UpdatedAt,
			&deletedAt,
		)
		if err != nil {
			return nil, err
		}
		if deletedAt.Valid {
			role.DeletedAt = deletedAt.Time
		}
		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *roleRepository) Update(ctx context.Context, role *model.Roles) error {
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
