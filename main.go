package main

import (
	"log"
	"os"

	"github.com/jimeh/ozu.io/shortner"
	"github.com/jimeh/ozu.io/storage/goleveldbstore"
	"github.com/jimeh/ozu.io/web"
	"github.com/valyala/fasthttp"
)

func main() {
	store, err := goleveldbstore.New("ozuio_database")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	shortner := shortner.New(store)
	router := web.NewRouter(shortner)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(fasthttp.ListenAndServe(":"+port, router.HandleRequest))
}
