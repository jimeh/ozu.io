package web

import (
	"encoding/json"
	"html/template"
	"mime"
	"net/url"
	"path"
	"time"

	"github.com/jimeh/ozu.io/shortener"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

//go:generate go-bindata -pkg web static/... templates/...

// NewHandler creates a new Handler object.
func NewHandler(s shortener.Shortener) *Handler {
	t := template.New("base")

	files, err := AssetDir("templates")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		content, err := Asset("templates/" + f)
		if err != nil {
			panic(err)
		}

		t.New(f).Parse(string(content))
	}

	return &Handler{s, t}
}

// Handler handle HTTP requests.
type Handler struct {
	s shortener.Shortener
	t *template.Template
}

// Index handles requests for root.
func (h *Handler) Index(c *routing.Context) error {
	h.template(c, "index.html", nil)
	return nil
}

// NotFound returns a 404 error page.
func (h *Handler) NotFound(c *routing.Context) error {
	c.NotFound()
	return nil
}

// Static returns assets serialized via go-bindata
func (h *Handler) Static(c *routing.Context) error {
	p := string(c.Path())[1:]

	data, err := Asset(p)
	if err != nil {
		h.NotFound(c)
		return nil
	}

	info, _ := AssetInfo(p)
	contentType := mime.TypeByExtension(path.Ext(p))
	modTime := info.ModTime().In(time.FixedZone("GMT", 0))

	c.SetContentType(contentType)
	c.Response.Header.Set("Last-Modified", modTime.Format(time.RFC1123))
	c.Write(data)
	return nil
}

// Shorten shortens given URL.
func (h *Handler) Shorten(c *routing.Context) error {
	uid, url, err := h.s.Shorten(c.FormValue("url"))
	if err != nil {
		return h.respondWithError(c, err)
	}

	r := h.makeURLResponse(c, uid, url)
	return h.respond(c, &r)
}

// Lookup shortened UID.
func (h *Handler) Lookup(c *routing.Context) error {
	uid := c.FormValue("uid")
	url, err := h.s.Lookup(uid)
	if err != nil {
		return h.respondWithError(c, err)
	}

	r := h.makeURLResponse(c, uid, url)
	return h.respond(c, &r)
}

// LookupAndRedirect looks up given UID and redirects to it's URL.
func (h *Handler) LookupAndRedirect(c *routing.Context) error {
	uid := []byte(c.Param("uid"))

	url, err := h.s.Lookup(uid)
	if err != nil {
		h.NotFound(c)
		return nil
	}

	r := h.makeURLResponse(c, uid, url)

	c.Response.Header.Set("Pragma", "no-cache")
	c.Response.Header.Set("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
	c.Response.Header.Set("X-XSS-Protection", "1; mode=block")
	c.Response.Header.Set("Cache-Control",
		"no-cache, no-store, max-age=0, must-revalidate")
	c.Redirect(string(url), fasthttp.StatusMovedPermanently)
	c.Response.Header.Set("Connection", "close")
	c.Response.Header.Set("X-Content-Type-Options", "nosniff")
	c.Response.Header.Set("Accept-Ranges", "none")
	c.Response.Header.Set("X-Frame-Options", "SAMEORIGIN")
	c.Response.Header.Set("Vary", "Accept-Encoding")

	h.template(c, "redirect.html", r)
	return nil
}

func (h *Handler) template(c *routing.Context, name string, data interface{}) {
	c.SetContentType("text/html; charset=UTF-8")
	h.t.ExecuteTemplate(c, name, data)
}

func (h *Handler) respond(c *routing.Context, r *URLResponse) error {
	resp, err := json.Marshal(r)
	if err != nil {
		return err
	}

	c.SetContentType("application/json")
	c.Write(resp)
	return nil
}

func (h *Handler) respondWithError(c *routing.Context, err error) error {
	r := ErrorResponse{
		Error: err.Error(),
	}

	resp, err := json.Marshal(r)
	if err != nil {
		return err
	}

	c.SetStatusCode(fasthttp.StatusNotFound)
	c.SetContentType("application/json")
	c.Write(resp)
	return nil
}

func (h *Handler) makeURLResponse(c *routing.Context, uid []byte, url []byte) URLResponse {
	return URLResponse{
		UID:    string(uid),
		URL:    h.makeShortURL(c, uid),
		Target: string(url),
	}
}

func (h *Handler) makeShortURL(c *routing.Context, uid []byte) string {
	shortURL := &url.URL{
		Scheme: "http",
		Host:   string(c.Host()),
		Path:   "/" + string(uid),
	}

	return shortURL.String()
}
