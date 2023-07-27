package repository

import (
	"context"
	"go-friend-sphere/internal/auth"
	"go-friend-sphere/internal/models"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type authRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(ctx context.Context, user *models.User) (*models.User, error) {
	return nil, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	foundedUser := &models.User{}
	if err := r.db.QueryRowxContext(ctx, findUserByEmail, user.Email).StructScan(foundedUser); err != nil {
		return nil, errors.Wrap(err, "findUserByEmail")
	}

	return foundedUser, nil
}
