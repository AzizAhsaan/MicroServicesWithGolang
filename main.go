package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"
	"os/signal"
	"github.com/AzizAhsaan/GolangMicroServices/handlers"
)

func main() {
	l := log.New(os.Stdout,"product-api",log.LstdFlags)
	products := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/",products)
	s := &http.Server{
		Addr :":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}
	go func() {

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	cha := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", cha)
	tc, _ := context.WithTimeout(context.Background(), 30 *time.Second)
	s.Shutdown(tc)
}