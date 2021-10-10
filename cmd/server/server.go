package main

import (
	"DZ_Backend_dev_Go_level_1/internal/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	dirToServe := "/tmp/upload"
	handler := &handlers.Handler{Subject: "Jon"}
	uploadHandler := &handlers.UploadHandler{
		HostAddr:  "http://localhost",
		UploadDir: dirToServe,
	}
	fileHendler := &handlers.FileHendler{
		PathDir: dirToServe,
	}

	// Файл-сервер можно закомментировать, вкладка files работает вместо него.
	http.Handle("/", http.FileServer(http.Dir(dirToServe)))

	http.Handle("/home", handler)
	http.Handle("/upload", uploadHandler)
	http.Handle("/files", fileHendler)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Println("Server started...")
		err := srv.ListenAndServe()
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
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("Erorr while shotting down the server: %v", err)
	}
	log.Println("Server stopped gracefully!")
}
