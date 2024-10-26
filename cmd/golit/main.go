package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/antomfdez/golit/internal/listener"
)

func main() {
	port := flag.String("p", "8080", "Listen Port")
	flag.Parse()
	if *port == "8080" {
		fmt.Printf("Using default port: %s\n", *port)
	} else {
		fmt.Printf("Listening on port: %s\n", *port)
	}
	server := listener.NewServer()
	if err := server.Listen(*port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
