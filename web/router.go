package web

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/jimeh/ozu.io/shortner"
)

func NewRouter(shortner *shortner.Shortner) *fasthttprouter.Router {
	router := fasthttprouter.New()

	return router
}
