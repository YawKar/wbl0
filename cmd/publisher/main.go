package main

import (
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
)

func main() {
	config := setAndParseFlagsIntoConfig()

	setupGlobalLogger(&config.LogConfig)
	slog.Debug("config:", config)

	sc, err := stan.Connect(
		config.NatsConfig.clusterId,
		config.NatsConfig.clientId,
		stan.NatsURL(config.NatsConfig.natsUrl),
	)
	if err != nil {
		log.Fatalln("Failed to connect to nats-streaming cluster!", err)
	}
	slog.Info("Successfully connected to Nats-streaming!")
	slog.Info("Settings setup for spam",
		"spam-duration", config.spamDuration,
		"spam-rate", config.spamRate,
		"seed", config.seed)

	generator := mkGen(config.seed)

	quit := time.After(config.spamDuration)
	ticker := time.NewTicker(config.spamRate)
FOR:
	for {
		select {
		case <-quit:
			slog.Info("Spam duration is over! Shutting down!")
			break FOR
		case <-ticker.C:
			if marshaled, err := proto.Marshal(generator()); err != nil {
				slog.Error("Couldn't Marshal fake message!", err)
			} else {
				sc.PublishAsync("orders", marshaled, ackHandler)
			}
		}
	}
	ticker.Stop()
}

func ackHandler(guid string, err error) {
	if err != nil {
		slog.Error("Ack error:", err)
	} else {
		slog.Info("Successfully acked", "guid", guid)
	}
}

func setupGlobalLogger(c *LogConfig) {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: &c.logLevelVar,
	}))
	slog.SetDefault(l)
}
