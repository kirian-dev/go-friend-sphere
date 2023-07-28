package repository

import (
	"context"
	"database/sql"
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
	u := &models.User{}
	if err := r.db.QueryRowxContext(ctx, createUser, user.Email, user.Password, user.FirstName, user.LastName).StructScan(u); err != nil {
		return nil, errors.Wrap(err, "auth repository, Register")
	}

	return u, nil
}

func (r *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	foundedUser := &models.User{}
	if err := r.db.QueryRowxContext(ctx, findUserByEmail, user.Email).StructScan(foundedUser); err != nil {
		return nil, errors.Wrap(err, "auth repository, FindUserByEmail")
	}

	return foundedUser, nil
}

func (r *authRepo) GetUsers(ctx context.Context) ([]*models.User, error) {
	usersList := []*models.User{}

	if err := r.db.SelectContext(ctx, &usersList, getUsers); err != nil {
		return nil, errors.Wrap(err, "auth repository, GetUsers")
	}

	return usersList, nil
}

func (r *authRepo) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	foundedUser := &models.User{}

	if err := r.db.QueryRowxContext(ctx, getUsersById, userId).StructScan(foundedUser); err != nil {
		return nil, errors.Wrap(err, "auth repository, GetUserById")
	}

	return foundedUser, nil
}

func (r *authRepo) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	updatedUser := &models.User{}

	if err := r.db.GetContext(ctx, updatedUser, updateUserQuery, &user.Email, &user.FirstName, &user.LastName, &user.UserID); err != nil {
		return nil, errors.Wrap(err, "auth repository, UpdateUser")
	}

	return updatedUser, nil
}

func (r *authRepo) DeleteUser(ctx context.Context, userId int64) error {
	result, err := r.db.ExecContext(ctx, deleteUserQuery, userId)
	if err != nil {
		return errors.Wrap(err, "auth repository, DeleteUser")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "auth repository, RowsAffected")
	}
	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "auth repository, rowsAffected")
	}

	return nil
}
