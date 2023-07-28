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
	existsUser, err := u.authRepo.FindByEmail(ctx, user)
	if existsUser != nil || err != nil {
		return nil, errors.New("User already exists")
	}

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
