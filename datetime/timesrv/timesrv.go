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

// connMap в эту мапу сохраняем подключения к серверу.
var connMap = make(map[*net.Conn]struct{})

func main() {

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	cfg := net.ListenConfig{
		KeepAlive: time.Minute,
	}
	l, err := cfg.Listen(ctx, "tcp", ":9001")
	if err != nil {
		log.Fatal(err)
	}
	wg := &sync.WaitGroup{}
	log.Println("im started!")

	// Запускаем горутину, которая будет считывать из консоли сервера сообщения,
	// перебирать все соединения в мапе и в каждое соединение отправлять сообщение.
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			for conn := range connMap {
				_, err = fmt.Fprintf(*conn, "now: %s\n", input.Text())
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			var conn net.Conn
			conn, err = l.Accept()
			if err != nil {
				log.Println(err)
			}
			if err == nil {
				connMap[&conn] = struct{}{}
				wg.Add(1)
				go handleConn(ctx, wg, conn)
			}
		}
	}()

	<-ctx.Done()

	log.Println("done")
	err = l.Close()
	if err != nil {
		log.Println(err)
	}
	wg.Wait()
	log.Println("exit")
}

func handleConn(ctx context.Context, wg *sync.WaitGroup, conn net.Conn) {
	defer wg.Done()
	defer func(conn net.Conn) {
		delete(connMap, &conn)
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	// Каждую секунду отправлять клиентам текущее время сервера.
	tck := time.NewTicker(time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case t := <-tck.C:
			_, err := fmt.Fprintf(conn, "now: %s\n", t)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
