package middleware

import (
	"go-friend-sphere/config"
	"go-friend-sphere/internal/auth"
	"go-friend-sphere/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	authUC auth.UseCase
	cfg    *config.Config
	logger logger.ZapLogger
}

// Middleware manager constructor
func NewMiddlewareManager(authUC auth.UseCase, cfg *config.Config, logger logger.ZapLogger) *MiddlewareManager {
	return &MiddlewareManager{authUC: authUC, cfg: cfg, logger: logger}
}
