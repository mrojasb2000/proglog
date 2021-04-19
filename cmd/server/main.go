package main

import (
	"log"

	"github.com/mrojasb2000/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":9090")
	log.Fatal(srv.ListenAndServe())
}
