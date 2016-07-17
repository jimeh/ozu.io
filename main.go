package main

import (
	"log"
	"os"

	"github.com/jimeh/ozu.io/shortener"
	"github.com/jimeh/ozu.io/storage/goleveldbstore"
	"github.com/jimeh/ozu.io/web"
	"github.com/syndtr/goleveldb/leveldb"
)

var path = "ozuio_database"

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return port
}

func main() {
	db, errDB := leveldb.OpenFile(path, nil)
	if errDB != nil {
		log.Fatal(errDB)
	}
	defer db.Close()

	store, err := goleveldbstore.New(db)
	if err != nil {
		log.Fatal(err)
	}

	s := shortener.New(store)
	server := web.NewServer(s)

	log.Fatal(server.ListenAndServe(":" + getPort()))
}
