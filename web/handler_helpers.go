package web

import (
	"net/url"

	"github.com/qiangxue/fasthttp-routing"
)

func makeResponse(c *routing.Context, uid []byte, url []byte) Response {
	return Response{
		UID:    string(uid),
		URL:    makeShortURL(c, uid),
		Target: string(url),
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
