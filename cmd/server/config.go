package main

import (
	"flag"
	"fmt"
	"log/slog"
	"time"

	"github.com/yawkar/wbl0/pkg/storage"
)

type Config struct {
	ServerConfig
	NatsConfig
	storage.StorageConfig
	LogConfig
}

type ServerConfig struct {
	addr string
	port uint
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
	flag.StringVar(&c.ServerConfig.addr, "addr", "localhost", "set server's address (server listens to and serves addr:port)")
	flag.UintVar(&c.ServerConfig.port, "port", 8080, "set server's port (server listens to and serves addr:port)")

	flag.StringVar(&c.NatsConfig.natsUrl, "nats-url", "nats://127.0.0.1:4222",
		"set nats-streaming node's url")
	flag.StringVar(&c.NatsConfig.clusterId, "cluster-id", "default_cluster",
		"set nats-streaming cluster's id")
	flag.StringVar(&c.NatsConfig.clientId, "client-id", "default_client",
		"set nats-streaming client's id")

	flag.StringVar(&c.StorageConfig.DbUrl, "db-url", "", "set database url")

	flag.DurationVar(&c.StorageConfig.CacheExpiration, "cache-expire", 3*time.Minute, "set cache expiration time")
	flag.DurationVar(&c.StorageConfig.CleanupInterval, "cache-cleanup", 6*time.Minute, "set cache cleanup interval")

	logLevel := flag.Int("log-level", 0, fmt.Sprintf(
		"debug: %d; info: %d; warn: %d; error: %d",
		slog.LevelDebug,
		slog.LevelInfo,
		slog.LevelWarn,
		slog.LevelError,
	))

	flag.Parse()
	c.LogConfig.logLevelVar.Set(slog.Level(*logLevel))
	return
}
