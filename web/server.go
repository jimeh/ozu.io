package web

import (
	"github.com/jimeh/ozu.io/shortener"
	"github.com/valyala/fasthttp"
)

// NewServer returns a new fasthttp.Server with all routes configured.
func NewServer(s shortener.Shortener) *fasthttp.Server {
	r := NewRouter(s)
	return &fasthttp.Server{Handler: r.HandleRequest}
}
