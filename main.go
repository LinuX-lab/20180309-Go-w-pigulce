//go:generate go-bindata-assetfs -pkg main -prefix www www

package main

import (
	"flag"
	"log"
)

var (
	endPoint = flag.String("address", ":8000", "Server address in form ADDR:PORT or just :PORT")
	sesDir   = flag.String("sessdir", "sessions", "Absolute or relative path for session info directory")
)

func main() {
	flag.Parse()

	// Wystartowanie huba pośredniczącego
	startHub()

	// Wystartowanie serera WWW
	err := startServer(*sesDir, *endPoint)
	if err != nil {
		log.Panicln("Error starting server: ", err)
	}

	// Zawieszenie głównego goroutine'a
	select {}
}
