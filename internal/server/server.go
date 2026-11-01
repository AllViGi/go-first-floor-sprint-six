package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type server struct {
	logger     *log.Logger
	HttpServer *http.Server
}

func NewServer(l *log.Logger) *server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandle)
	mux.HandleFunc("/upload", handlers.UploadHandle)

	myServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     l,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}
	return &server{logger: l, HttpServer: myServer}
}
