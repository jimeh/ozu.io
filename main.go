package main

import (
	"log"
	"os"

	"github.com/jimeh/ozu.io/shortener"
	"github.com/jimeh/ozu.io/storage/goleveldbstore"
	"github.com/jimeh/ozu.io/web"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}

func main() {
	store, err := goleveldbstore.New("ozuio_database")
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	s := shortener.New(store)
	server := web.NewServer(s)

	log.Fatal(server.ListenAndServe(":" + getPort()))
}
