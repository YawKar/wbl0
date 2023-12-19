package main

import (
	"database/sql"
	"log"
	"log/slog"
	"os"

	"github.com/YawKar/wbl0/pkg/storage"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
)

func main() {
	config := setAndParseFlagsIntoConfig()

	// setup logger
	setupGlobalLogger(&config.LogConfig)
	slog.Debug("config:", config)

	// create db connection
	db, err := storage.MkDb(&config.DbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// create cache
	cache, err := storage.MkCache(&config.CacheConfig)
	if err != nil {
		log.Fatalln(err)
	}

	// populate cache from db
	// TODO

	// connect to nats-streaming cluster and subscribe
	sc, err := stan.Connect(
		config.NatsConfig.clusterId,
		config.NatsConfig.clientId,
		stan.NatsURL(config.NatsConfig.natsUrl),
	)
	if err != nil {
		log.Fatalln("Failed to connect to nats-streaming cluster!", err)
	}
	sc.Subscribe("orders", mkMsgHandler(db, cache), stan.SetManualAckMode())
}

func mkMsgHandler(db *sql.DB, cache *cache.Cache) func(*stan.Msg) {
	return func(m *stan.Msg) {
		slog.Debug("received a message:", m)
		log.Printf("message: %s\n", string(m.Data))
	}
}

func setupGlobalLogger(c *LogConfig) {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: &c.logLevelVar,
	}))
	slog.SetDefault(l)
}
