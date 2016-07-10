package main

import (
	"fmt"
	"log"

	"github.com/jimeh/ozu.io/shortner"
	"github.com/jimeh/ozu.io/storage/goleveldbstore"
	"github.com/jimeh/ozu.io/web"
)

func main() {
	store, err := goleveldbstore.New("ozuio_database")
	if err != nil {
		log.Fatal(err)
	}

	shortner := shortner.New(store)
	router := web.NewRouter(shortner)
	fmt.Println(router)
}
