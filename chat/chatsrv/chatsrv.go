package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	ch := make(chan string)
	go clientWriter(conn, ch)

	fmt.Fprintln(conn, "Введите свой никнейм:")
	input := bufio.NewScanner(conn)
	input.Scan()
	nik := input.Text()

	who := conn.RemoteAddr().String()
	if nik != "" {
		who = fmt.Sprintf("[ %s ]", nik)
	}

	// Выводит в консоль подключившегося пользователя сообщение как он зарегистрировался на сервере.
	ch <- fmt.Sprintf("You are %s", who)
	// Сообщение отправляется всем пользователям, о подключении нового пользователя.
	messages <- fmt.Sprintf("%s: has arrived", who)
	entering <- ch
	// Выводит в консоль сервера сообщение о подключении нового пользователя.
	log.Printf( "%s has arrived", who)

	//input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- fmt.Sprintf("%s: %s", who, input.Text())
	}
	leaving <- ch
	messages <- fmt.Sprintf( "%s: has left", who)
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}