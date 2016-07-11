package web

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/jimeh/ozu.io/shortner"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

// Handlers handle HTTP requests.
type Handlers struct {
	s *shortner.Shortner
}

// Index handles requests for root.
func (h *Handlers) Index(c *routing.Context) error {
	c.WriteString("Welcome to ozu.io, a shitty URL shortner.")
	return nil
}

// Shorten shortens given URL.
func (h *Handlers) Shorten(c *routing.Context) error {
	c.SetContentType("application/json")

	uid, url, err := h.s.Shorten(c.FormValue("url"))
	if err != nil {
		c.SetStatusCode(fasthttp.StatusBadRequest)
		response, _ := json.Marshal(ErrorResponse{err.Error()})
		c.Write(response)
		return nil
	}

	h.respondWithShortened(c, uid, url)
	return nil
}

// Lookup shortened UID.
func (h *Handlers) Lookup(c *routing.Context) error {
	c.SetContentType("application/json")

	uid := c.FormValue("uid")
	url, err := h.s.Lookup(uid)
	if err != nil {
		c.SetStatusCode(fasthttp.StatusNotFound)
		respBytes, _ := json.Marshal(ErrorResponse{err.Error()})
		c.Write(respBytes)
		return nil
	}

	h.respondWithShortened(c, uid, url)
	return nil
}

// LookupAndRedirect looks up given UID and redirects to it's URL.
func (h *Handlers) LookupAndRedirect(c *routing.Context) error {
	uid := c.Param("uid")
	url, err := h.s.Lookup([]byte(uid))
	if err != nil {
		c.SetStatusCode(fasthttp.StatusNotFound)
		fmt.Fprint(c, "404 Not Found")
		return nil
	}

	c.Redirect(string(url), fasthttp.StatusMovedPermanently)
	return nil
}

func (h *Handlers) respondWithShortened(c *routing.Context, uid []byte, url []byte) {
	c.SetStatusCode(fasthttp.StatusOK)
	response := ShortenedResponse{
		UID:      string(uid),
		ShortURL: h.makeShortURL(c, uid),
		URL:      string(url),
	}
	respBytes, _ := json.Marshal(response)
	c.Write(respBytes)
}

func (h *Handlers) makeShortURL(c *routing.Context, uid []byte) string {
	shortURL := &url.URL{
		Scheme: "http",
		Host:   string(c.Host()),
		Path:   "/" + string(uid),
	}

	return shortURL.String()
}
