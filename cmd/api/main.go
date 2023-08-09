package main

import (
	config "go-friend-sphere/config"
	"go-friend-sphere/internal/server"
	"go-friend-sphere/pkg/db/postgres"
	"go-friend-sphere/pkg/helpers"
	"go-friend-sphere/pkg/logger"
	"log"
	"os"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// @title GO-Friend-Sphere
// @version 1.0
// @description Pet-project GO-Friend-Sphere REST API
// @contact.email polozenko.kirill.job@gmail.com
// @BasePath /
// @host localhost:8080

func main() {
	log.Println("Starting server")

	configPath := helpers.GetConfigPath(os.Getenv("config"))

	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	logger := logger.InitLogger(cfg)
	logger.Infof("App version: %s, Mode: %s", cfg.Server.AppVersion, cfg.Server.Mode)

	psqlDB, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Fatalf("Failed to create database: %v", err)
	} else {
		logger.Infof("Postgres connected, status: %v", psqlDB.Stats())
	}

	defer psqlDB.Close()

	jaegerCfgInstance := jaegercfg.Configuration{
		ServiceName: cfg.Jaeger.ServiceName,
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           cfg.Jaeger.LogSpans,
			LocalAgentHostPort: cfg.Jaeger.Host,
		},
	}

	tracer, closer, err := jaegerCfgInstance.NewTracer(
		jaegercfg.Logger(jaegerlog.StdLogger),
		jaegercfg.Metrics(metrics.NullFactory),
	)
	if err != nil {
		log.Fatal("could not create a tracer", err)
	}
	logger.Info("Jaeger connected successfully")

	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()
	logger.Info("Opentracing connected successfully")

	s := server.NewServer(cfg, *logger, psqlDB)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
