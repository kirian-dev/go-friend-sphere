package usecase

import (
	"context"
	"go-friend-sphere/config"
	"go-friend-sphere/internal/auth"
	"go-friend-sphere/internal/models"
	"go-friend-sphere/pkg/logger"

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
		return nil, errors.New("User with given email already exists")
	}

	return existsUser, nil
}
