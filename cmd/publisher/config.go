package main

import (
	"flag"
	"fmt"
	"log/slog"
	"time"
)

type Config struct {
	PublisherConfig
	NatsConfig
	LogConfig
}

type PublisherConfig struct {
	spamDuration time.Duration
	spamRate     time.Duration
	seed         int64
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
	flag.DurationVar(&c.spamDuration, "spam-duration", 1*time.Minute, "set spam duration")
	flag.DurationVar(&c.spamRate, "spam-rate", 1*time.Second, "set spam rate (2s means 1 message every 2 seconds)")
	flag.Int64Var(&c.seed, "seed", 42, "set seed for faker")

	flag.StringVar(&c.NatsConfig.natsUrl, "nats-url", "nats://127.0.0.1:4222",
		"set nats-streaming node's url")
	flag.StringVar(&c.NatsConfig.clusterId, "cluster-id", "default_cluster",
		"set nats-streaming cluster's id")
	flag.StringVar(&c.NatsConfig.clientId, "client-id", "default_client",
		"set nats-streaming client's id")

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
