package main

import (
	"flag"
	"fmt"
	"log/slog"
	"time"

	"github.com/YawKar/wbl0/pkg/storage"
)

type Config struct {
	NatsConfig
	storage.DbConfig
	storage.CacheConfig
	LogConfig
}

type NatsConfig struct {
	natsUrl   string
	clusterId string
	clientId  string
}

type LogConfig struct {
	logLevelVar slog.LevelVar
}

func setAndParseFlagsIntoConfig() (c Config) {
	flag.StringVar(&c.NatsConfig.natsUrl, "nats-url", "nats://127.0.0.1:4222",
		"set nats-streaming node's url")
	flag.StringVar(&c.NatsConfig.clusterId, "cluster-id", "default_cluster",
		"set nats-streaming cluster's id")
	flag.StringVar(&c.NatsConfig.clientId, "client-id", "default_client",
		"set nats-streaming client's id")

	flag.StringVar(&c.DbConfig.DbUrl, "db-url", "", "set database url")

	flag.DurationVar(&c.CacheConfig.CacheExpiration, "cache-expire", 3*time.Minute, "set cache expiration time")
	flag.DurationVar(&c.CacheConfig.CleanupInterval, "cache-cleanup", 6*time.Minute, "set cache cleanup interval")

	logLevel := flag.Int("log-level", 0, fmt.Sprintf(
		"debug: %d; info: %d; warn: %d; error: %d",
		slog.LevelInfo,
		slog.LevelDebug,
		slog.LevelWarn,
		slog.LevelError,
	))

	flag.Parse()
	c.LogConfig.logLevelVar.Set(slog.Level(*logLevel))
	return
}
