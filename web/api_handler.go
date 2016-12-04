package web

import (
	"encoding/json"

	"github.com/jimeh/ozu.io/shortener"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// NewAPIHandler creates a new Handler object.
func NewAPIHandler(s shortener.Shortener) *APIHandler {
	return &APIHandler{s}
}

// APIHandler handle HTTP requests.
type APIHandler struct {
	shortener shortener.Shortener
}

// Shorten shortens given URL.
func (h *APIHandler) Shorten(c *routing.Context) error {
	record, err := h.shortener.Shorten(c.FormValue("url"))
	if err != nil {
		return h.respondWithError(c, err)
	}

	r := makeResponse(c, record)
	return h.respond(c, &r)
}

// Lookup shortened UID.
func (h *APIHandler) Lookup(c *routing.Context) error {
	uid := c.FormValue("uid")
	record, err := h.shortener.Lookup(uid)
	if err != nil {
		return h.respondWithError(c, err)
	}

	r := makeResponse(c, record)
	return h.respond(c, &r)
}

func (h *APIHandler) respond(c *routing.Context, r *Response) error {
	resp, err := json.Marshal(r)
	if err != nil {
		return err
	}

	c.SetContentType("application/json")
	c.Write(resp)
	return nil
}

func (h *APIHandler) respondWithError(c *routing.Context, err error) error {
	r := Response{Error: err.Error()}
	resp, err := json.Marshal(r)
	if err != nil {
		return err
	}

	c.SetStatusCode(fasthttp.StatusNotFound)
	c.SetContentType("application/json")
	c.Write(resp)
	return nil
}
