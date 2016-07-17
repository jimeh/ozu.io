package web

import (
	"github.com/jimeh/ozu.io/shortener"
	"github.com/qiangxue/fasthttp-routing"
)

// NewRouter creates a new routing.Router with all handlers registered.
func NewRouter(s shortener.Shortener) *routing.Router {
	r := routing.New()
	h := NewHandler(s)

	r.Get("/", h.Index)
	r.Get("/api/shorten", h.Shorten)
	r.Get("/api/lookup", h.Lookup)
	r.Get("/static/*", h.Static)
	r.Get("/<uid>", h.LookupAndRedirect)
	r.Get("/*", h.NotFound)

	return r
}
