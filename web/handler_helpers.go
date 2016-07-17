package web

import (
	"net/url"

	"github.com/qiangxue/fasthttp-routing"
)

func makeURLResponse(c *routing.Context, uid []byte, url []byte) URLResponse {
	return URLResponse{
		UID:    string(uid),
		URL:    makeShortURL(c, uid),
		Target: string(url),
	}
}

func makeShortURL(c *routing.Context, uid []byte) string {
	shortURL := &url.URL{
		Scheme: "http",
		Host:   string(c.Host()),
		Path:   "/" + string(uid),
	}

	return shortURL.String()
}
