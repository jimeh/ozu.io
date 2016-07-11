package web

import (
	"github.com/jimeh/ozu.io/shortener"
	"github.com/qiangxue/fasthttp-routing"
)

// NewRouter creates a new routing.Router with all handlers registered.
func NewRouter(shortener *shortener.Shortener) *routing.Router {
	router := routing.New()
	handlers := Handlers{shortener}

	router.Get("/", handlers.Index)
	router.Get("/api/shorten", handlers.Shorten)
	router.Get("/api/lookup", handlers.Lookup)
	router.Get("/<uid>", handlers.LookupAndRedirect)

	return router
}
