package main

import (
	"embed"
	"html/template"

	"github.com/yawkar/wbl0/pkg/models"
)

//go:embed html
var html embed.FS

type ViewPageData struct {
	models.Order
	models.Payment
	models.Delivery
	Items []models.Item
}

func getBasicViewTemplate() (*template.Template, error) {
	return template.ParseFS(html, "html/basicview.html")
}
