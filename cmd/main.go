package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)

	newServer := server.NewServer(logger)

	if err := http.ListenAndServe(newServer.HttpServer.Addr, newServer.HttpServer.Handler); err != nil {
		log.Fatal(err)
	}

}
