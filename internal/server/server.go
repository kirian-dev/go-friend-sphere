package server

import (
	"context"
	config "go-friend-sphere/conifg"
	"go-friend-sphere/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	maxHeadersBytes = 1 << 20
	ctxTimeout      = 5
)

type Server struct {
	cfg    *config.Config
	logger *logger.ZapLogger
	db     *sqlx.DB
}

func NewServer(cfg *config.Config, logger *logger.ZapLogger, db *sqlx.DB) *Server {
	return &Server{cfg: cfg, logger: logger, db: db}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:           s.cfg.Server.Port,
		ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
		WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
		MaxHeaderBytes: maxHeadersBytes,
	}

	go func() {
		s.logger.Info("Server is listening on PORT %s", s.cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatalf("Error starting server: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		s.logger.Fatalf("Error shutting down server: %v", err)
	}

	s.logger.Info("Server Exited Properly")
	return nil
}
