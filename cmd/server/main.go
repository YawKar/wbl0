package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/nats-io/stan.go"
	pb "github.com/yawkar/wbl0/pkg/proto"
	"github.com/yawkar/wbl0/pkg/storage"
	"google.golang.org/protobuf/proto"
)

func main() {
	config := setAndParseFlagsIntoConfig()

	// setup logger
	setupGlobalLogger(&config.LogConfig)
	slog.Debug("config:", config)

	// make storage
	store, err := storage.MkStorage(&config.StorageConfig, false)
	if err != nil {
		log.Fatalln(err)
	}

	// connect to nats-streaming cluster and subscribe
	sc, err := stan.Connect(
		config.NatsConfig.clusterId,
		config.NatsConfig.clientId,
		stan.NatsURL(config.NatsConfig.natsUrl),
	)
	if err != nil {
		log.Fatalln("Failed to connect to nats-streaming cluster!", err)
	}
	sc.Subscribe("orders", mkMsgHandler(store))

	// dummy wait
	boom := time.After(time.Hour)
	for range boom {
		fmt.Println("That's all folks!")
		return
	}
}

func mkMsgHandler(store *storage.Storage) func(*stan.Msg) {
	return func(m *stan.Msg) {
		slog.Debug("received a nats message:", "message", m)
		order := &pb.Order{}
		if err := proto.Unmarshal(m.Data, order); err != nil {
			slog.Error("Couldn't unmarshal message!", "error", err)
		} else if err := deepInsertOrder(store, order); err != nil {
			slog.Error("Couldn't insert order!", "error", err)
		} else {
			slog.Debug("Successfully inserted order", "uuid", order.OrderUid)
		}
	}
}

func setupGlobalLogger(c *LogConfig) {
	l := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: &c.logLevelVar,
	}))
	slog.SetDefault(l)
}
