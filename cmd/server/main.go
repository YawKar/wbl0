package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/google/uuid"
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

	// setup orderRes env
	orderRes := OrderResource{store: store}

	// setup gin server
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Mount("/order", orderRes.Routes())
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		log.Fatalln("server failed to serve", "err", err)
	}
}

type OrderResource struct {
	store *storage.Storage
}

func (o *OrderResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Route(
		"/{uuid:(?i)[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}}",
		func(r chi.Router) {
			r.Get("/", o.hGetOrder)
			r.Get("/payment", o.hGetOrderPayment)
			r.Get("/delivery", o.hGetOrderDelivery)
			r.Get("/items", o.hGetOrderItems)
		},
	)
	return r
}

func (h *OrderResource) hGetOrder(w http.ResponseWriter, r *http.Request) {
	slog.Info("gotcha")
	uuidParam := chi.URLParam(r, "uuid")
	if uuidParam == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		http.Error(w, "provided uuid isn't correct uuid", http.StatusBadRequest)
		return
	}
	order, err := h.store.GetOrder(uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("order with uuid = %s wasn't found", uuid), http.StatusNotFound)
		return
	}
	if data, err := json.Marshal(order); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Write(data)
	}
}

func (h *OrderResource) hGetOrderPayment(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")
	if uuidParam == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		http.Error(w, "provided uuid isn't correct uuid", http.StatusBadRequest)
		return
	}
	payment, err := h.store.GetPayment(uuid);
	if err != nil {
		http.Error(w, fmt.Sprintf("payment details for order with uuid = %s weren't found", uuid), http.StatusNotFound)
		return
	}
	if data, err := json.Marshal(payment); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Write(data)
	}
}

func (h *OrderResource) hGetOrderDelivery(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")
	if uuidParam == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		http.Error(w, "provided uuid isn't correct uuid", http.StatusBadRequest)
		return
	}
	delivery, err := h.store.GetDelivery(uuid);
	if err != nil {
		http.Error(w, fmt.Sprintf("delivery details for order with uuid = %s weren't found", uuid), http.StatusNotFound)
		return
	}
	if data, err := json.Marshal(delivery); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Write(data)
	}
}

func (h *OrderResource) hGetOrderItems(w http.ResponseWriter, r *http.Request) {
	uuidParam := chi.URLParam(r, "uuid")
	if uuidParam == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	uuid, err := uuid.Parse(uuidParam)
	if err != nil {
		http.Error(w, "provided uuid isn't correct uuid", http.StatusBadRequest)
		return
	}
	items, err := h.store.GetItems(uuid);
	if err != nil {
		http.Error(w, fmt.Sprintf("items for order with uuid = %s weren't found", uuid), http.StatusNotFound)
		return
	}
	if data, err := json.Marshal(items); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		w.Write(data)
	}
}

func mkMsgHandler(store *storage.Storage) func(*stan.Msg) {
	return func(m *stan.Msg) {
		slog.Debug("received a nats message:", "message", m)
		order := &pb.Order{}
		if err := proto.Unmarshal(m.Data, order); err != nil {
			slog.Error("Couldn't unmarshal message!", "error", err)
		} else if err := insertOrderMessage(store, order); err != nil {
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
