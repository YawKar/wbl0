package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/yawkar/wbl0/pkg/models"
	"github.com/yawkar/wbl0/pkg/storage"
)

type viewResource struct {
	store             *storage.Storage
	viewOrderTemplate *template.Template
}

func NewViewResource(store *storage.Storage) (HandlersResource, error) {
	viewOrder, err := getBasicViewTemplate()
	if err != nil {
		return nil, err
	}
	return &viewResource{store: store, viewOrderTemplate: viewOrder}, nil
}

func (o *viewResource) Routes() chi.Router {
	r := chi.NewRouter()
	r.Get("/{uuid:(?i)[0-9A-F]{8}-[0-9A-F]{4}-[4][0-9A-F]{3}-[89AB][0-9A-F]{3}-[0-9A-F]{12}}", o.hViewOrder)
	return r
}

func (h *viewResource) hViewOrder(w http.ResponseWriter, r *http.Request) {
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
	payment, err := h.store.GetPayment(uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("payment for order with uuid = %s wasn't found", uuid), http.StatusNotFound)
		return
	}
	delivery, err := h.store.GetDelivery(uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("delivery for order with uuid = %s wasn't found", uuid), http.StatusNotFound)
		return
	}
	items, err := h.store.GetItems(uuid)
	if err != nil {
		http.Error(w, fmt.Sprintf("items for order with uuid = %s weren't found", uuid), http.StatusNotFound)
		return
	}
	itemsV := make([]models.Item, 0)
	for _, item := range items {
		itemsV = append(itemsV, *item)
	}
	basicViewData := ViewPageData{
		Order:    *order,
		Payment:  *payment,
		Delivery: *delivery,
		Items:    itemsV,
	}
	h.viewOrderTemplate.Execute(w, basicViewData)
}
