package web

import (
	"github.com/jimeh/ozu.io/shortener"
	"github.com/qiangxue/fasthttp-routing"
)

// NewRouter creates a new routing.Router with all handlers registered.
func NewRouter(s shortener.Shortener) *routing.Router {
	r := routing.New()

	api := NewAPIHandler(s)
	r.Get("/api/shorten", api.Shorten)
	r.Get("/api/lookup", api.Lookup)

	handler := NewHandler(s)
	r.Get("/", handler.Index)
	r.Get("/static/*", handler.Static)
	r.Get("/<uid>", handler.LookupAndRedirect)
	r.Get("/*", handler.NotFound)

	return r
}
