package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	cfg := net.ListenConfig{
		KeepAlive: time.Minute,
	}
	l, err := cfg.Listen(ctx, "tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	log.Println("im started!")

	// countConn в эту переменную сохраняем количество подключений к серверу.
	countConn := 0
	// Создаём новый канал для отправки сообщений.
	message := make(chan string)
	// Запускаем горутину, которая будет считывать из консоли сервера сообщения,
	// и отправлять в канал message.
	go func(int) {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			for i := 0; i < countConn; i++ {
				message <- "timesrv msg: " + input.Text()
			}
		}
	}(countConn)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			conn, err := l.Accept()
			if err != nil {
				log.Println(err)
			}
			if err == nil {
				countConn++
				wg.Add(1)
				go handleConn(ctx, wg, conn, message)
			}
		}
	}()

	<-ctx.Done()

	log.Println("done")
	l.Close()
	wg.Wait()
	log.Println("exit")
}

func handleConn(ctx context.Context, wg *sync.WaitGroup, conn net.Conn, message chan string) {
	defer wg.Done()
	defer conn.Close()

	// Каждую секунду отправлять клиентам текущее время сервера.
	tck := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case t := <-tck.C:
			fmt.Fprintf(conn, "now: %s\n", t)
		case msg := <-message:
			fmt.Fprintf(conn, "now: %s\n", msg)
		}
	}
}
