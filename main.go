package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func LogServer(w http.ResponseWriter, req *http.Request) {
	log.Println("", req.RequestURI)
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", LogServer)
	srv := &http.Server{
		Addr:         "0.0.0.0:80",
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		Handler:      r,
	}
	log.Println("Initialization complete, starting server.")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println("Stopped serving", "err", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down")
}
