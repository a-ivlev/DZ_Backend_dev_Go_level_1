package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func main() {
	dirToServe := http.Dir("/tmp/upload")
	//uploadHandler := &UploadHandler{
	//	dirToServe: http.Dir("/tmp/upload"),
	//}

	//s := http.StripPrefix("/upload/", http.FileServer(dirToServe))
	//http.Handle("/upload", s)
	http.Handle("/upload", http.StripPrefix("/upload", http.FileServer(dirToServe)))
	//http.Handle("/upload", uploadHandler)
	//http.Handle("/get/", http.FileServer(httpDir))

	fs := &http.Server{
		Addr: ":8081",
		//Handler:      http.FileServer(dirToServe),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Server started...")
		err := fs.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server returned an error: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	switch <-interrupt {
	case os.Interrupt:
		log.Println("Got SIGINT or SIGKILL...")
	case syscall.SIGTERM:
		log.Println("Got SIGTERM...")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err := fs.Shutdown(ctx)
	if err != nil {
		log.Printf("Error while shutting down the server: %v", err)
	}
	log.Println("Server stopped gracefully!")
}

