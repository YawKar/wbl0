package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/yawkar/wbl0/pkg/storage"
)

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
	payment, err := h.store.GetPayment(uuid)
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
	delivery, err := h.store.GetDelivery(uuid)
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
	items, err := h.store.GetItems(uuid)
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