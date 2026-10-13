package branch

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var ErrNotFound = errors.New("branch not found")

type Branch struct {
	ID       string `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Phone    string `json:"phone,omitempty"`
	Email    string `json:"email,omitempty"`
	Address  string `json:"address,omitempty"`
	Timezone string `json:"timezone"`
}

type Lister interface {
	List(ctx context.Context) ([]Branch, error)
}

type Finder interface {
	FindByID(ctx context.Context, id string) (Branch, error)
}

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List(ctx context.Context) ([]Branch, error) {
	const query = `
		SELECT
			id::text,
			code,
			name,
			COALESCE(phone, ''),
			COALESCE(email, ''),
			COALESCE(address, ''),
			timezone
		FROM branches
		WHERE status = 'ACTIVE'
		ORDER BY name, id
	`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list active branches: %w", err)
	}
	defer rows.Close()

	branches := make([]Branch, 0)

	for rows.Next() {
		var item Branch

		err := rows.Scan(
			&item.ID,
			&item.Code,
			&item.Name,
			&item.Phone,
			&item.Email,
			&item.Address,
			&item.Timezone,
		)
		if err != nil {
			return nil, fmt.Errorf("scan branch: %w", err)
		}

		branches = append(branches, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate branches: %w", err)
	}

	return branches, nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (Branch, error) {
	const query = `
		SELECT
			id::text,
			code,
			name,
			COALESCE(phone, ''),
			COALESCE(email, ''),
			COALESCE(address, ''),
			timezone
		FROM branches
		WHERE id = $1 AND status = 'ACTIVE'
	`

	var branch Branch
	err := r.db.QueryRow(ctx, query, id).Scan(
		&branch.ID,
		&branch.Code,
		&branch.Name,
		&branch.Phone,
		&branch.Email,
		&branch.Address,
		&branch.Timezone,
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return Branch{}, ErrNotFound
	}
	if err != nil {
		return Branch{}, fmt.Errorf("find active branch by ID: %w", err)
	}

	return branch, nil
}
