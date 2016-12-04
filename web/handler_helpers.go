package web

import (
	"net/url"

	"github.com/jimeh/ozu.io/storage"
	"github.com/qiangxue/fasthttp-routing"
)

func makeResponse(c *routing.Context, r *storage.Record) Response {
	return Response{
		UID:    string(r.UID),
		URL:    makeShortURL(c, r.UID),
		Target: string(r.URL),
	}
}

func makeErrResponse(err error) Response {
	return Response{Error: err.Error()}
}

func makeShortURL(c *routing.Context, uid []byte) string {
	shortURL := &url.URL{
		Scheme: "http",
		Host:   string(c.Host()),
		Path:   "/" + string(uid),
	}

	return shortURL.String()
}
