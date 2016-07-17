package web

import (
	"html/template"
	"mime"
	"path"
	"time"

	"github.com/jimeh/ozu.io/shortener"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
)

//go:generate go-bindata -pkg web static/... templates/...

// NewHandler creates a new Handler object.
func NewHandler(s shortener.Shortener) *Handler {
	t := newHandlerTemplate()
	return &Handler{s, t}
}

func newHandlerTemplate() *template.Template {
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

	return t
}

// Handler handle HTTP requests.
type Handler struct {
	shortener shortener.Shortener
	template  *template.Template
}

// Index handles requests for root.
func (h *Handler) Index(c *routing.Context) error {
	h.respond(c, "index.html", nil)
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

// LookupAndRedirect looks up given UID and redirects to it's URL.
func (h *Handler) LookupAndRedirect(c *routing.Context) error {
	uid := []byte(c.Param("uid"))

	url, err := h.shortener.Lookup(uid)
	if err != nil {
		h.NotFound(c)
		return nil
	}

	r := makeURLResponse(c, uid, url)

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

	h.respond(c, "redirect.html", r)
	return nil
}

func (h *Handler) respond(c *routing.Context, name string, data interface{}) {
	c.SetContentType("text/html; charset=UTF-8")
	h.template.ExecuteTemplate(c, name, data)
}
