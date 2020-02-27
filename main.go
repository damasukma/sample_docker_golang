package main

import (
	"context"
	"core_api/controller"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/", controller.HelloWorld)

	server := &http.Server{
		Addr: "0.0.0.0:9000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      route, // Pass our instance of gorilla/mux in.
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ShutDown(server)

}

func ShutDown(serv *http.Server) {
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	serv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)

}
