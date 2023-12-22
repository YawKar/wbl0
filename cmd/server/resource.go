package main

import "github.com/go-chi/chi"

type HandlersResource interface {
	Routes() chi.Router
}
