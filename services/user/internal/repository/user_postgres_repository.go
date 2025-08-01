package repository

import (
	"context"
	"database/sql"

	"github.com/mateusneves122/social-media/services/user/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) UserRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {

	query := `
		INSERT INTO user.users (name, email, password)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil
}
