package main

import (
	"os"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

func main() {
	shortner := NewShortner()
	defer shortner.Close()

	routes := Routes{Shortner: shortner}

	iris.Use(logger.New(iris.Logger))
	iris.Get("/set/:key/:value", routes.Set)
	iris.Get("/get/:key", routes.Get)
	iris.Get("/shorten/:url", routes.Shorten)
	iris.Get("/lookup/:uid", routes.Lookup)
	iris.Get("/", routes.Root)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	iris.Listen(":" + port)
}
