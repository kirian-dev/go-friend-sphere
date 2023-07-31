package usecase

import (
	"context"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/auth"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"strings"

	"github.com/pkg/errors"
)

type authUC struct {
	cfg      *config.Config
	authRepo auth.Repository
	logger   logger.ZapLogger
}

func NewAuthUC(cfg *config.Config, authRepo auth.Repository, logger logger.ZapLogger) auth.UseCase {
	return &authUC{cfg: cfg, authRepo: authRepo, logger: logger}
}

func (u *authUC) Register(ctx context.Context, user *models.User) (*models.User, error) {
	// existsUser, err := u.authRepo.FindByEmail(ctx, user)
	// if existsUser != nil || err != nil {
	// 	return nil, errors.New("Invalid credentials")
	// }

	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	user.Password = strings.TrimSpace(user.Password)
	if err := helpers.HashPassword(user); err != nil {
		return nil, errors.New("Failed to hash password")
	}

	createdUser, err := u.authRepo.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	helpers.RemovePassword(createdUser)

	return createdUser, nil
}

func (u *authUC) Login(ctx context.Context, user *models.User) (*models.User, error) {
	foundUser, err := u.authRepo.FindByEmail(ctx, user)
	if err != nil {
		return nil, err
	}
	if err = helpers.ComparePasswords(foundUser, user.Password); err != nil {
		return nil, errors.New("Invalid credentials")
	}

	helpers.RemovePassword(foundUser)
	return foundUser, nil
}

func (u *authUC) GetUsers(ctx context.Context) ([]*models.User, error) {
	return u.authRepo.GetUsers(ctx)
}

func (u *authUC) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	return u.authRepo.GetUserById(ctx, userId)
}

func (u *authUC) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))
	if user.Phone != nil {
		*user.Phone = strings.TrimSpace(*user.Phone)
	}
	updatedUser, err := u.authRepo.UpdateUser(ctx, user)
	if err != nil {
		return nil, errors.New("Error updating user: " + err.Error())
	}

	helpers.RemovePassword(updatedUser)

	return updatedUser, nil
}

func (u *authUC) DeleteUser(ctx context.Context, userId int64) error {
	return u.authRepo.DeleteUser(ctx, userId)
}
