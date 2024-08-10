package main

import (
	"flag"
	"log"

	"github.com/kuzin57/grpc-example/services/gateway/cmd/app"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "port to run")
	flag.Parse()

	appServer := app.NewServer(port)

	if err := appServer.Init(); err != nil {
		log.Fatal(err)
	}

	defer appServer.Stop()

	if err := appServer.Run(); err != nil {
		log.Fatal(err)
	}
}
