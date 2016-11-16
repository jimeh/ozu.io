package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jimeh/ozu.io/shortener"
	"github.com/jimeh/ozu.io/storage/goleveldbstore"
	"github.com/jimeh/ozu.io/web"
	"gopkg.in/alecthomas/kingpin.v2"
)

// Version gets populated with version at build-time.
var Version string
var defaultPort = "8080"

var (
	port = kingpin.Flag("port", "Port to listen to.").Short('p').
		Default(defaultPort).String()
	bind = kingpin.Flag("bind", "Bind address.").Short('b').
		Default("0.0.0.0").String()
	dir = kingpin.Flag("dir", "Directory to store database file.").
		Short('d').Default("ozuio_database").String()
	version = kingpin.Flag("version", "Print version info.").
		Short('v').Bool()
)

func printVersion() {
	fmt.Println("ozuio " + Version)
}

func startServer() {
	store, err := goleveldbstore.New(*dir)
	if err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	s := shortener.New(store)
	server := web.NewServer(s)

	if *port == defaultPort {
		envPort := os.Getenv("PORT")
		if envPort != "" {
			*port = envPort
		}
	}

	address := *bind + ":" + *port
	fmt.Println("Listening on " + address)
	log.Fatal(server.ListenAndServe(address))
}

func main() {
	kingpin.Parse()

	if *version {
		printVersion()
	} else {
		startServer()
	}
}
