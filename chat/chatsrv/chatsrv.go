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
	log.Println("Server is started!")

	go broadcaster()
	var conn net.Conn
	for {
		conn, err = listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	_, err := fmt.Fprintln(conn, "Для отключения от сервера и завершения сеанса нужно набрать EXIT")
	if err != nil {
		log.Println(err)
	}
	_, err = fmt.Fprintln(conn, "Введите свой никнейм:")
	if err != nil {
		log.Println(err)
	}

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
	log.Printf("%s has arrived", who)

	for input.Scan() {
		switch input.Text() {
		case "EXIT":
			break
		case "new nikname":
			//leaving <- ch
			messages <- fmt.Sprintf("%s: has left", who)
			log.Printf("%s: has left", who)
			input.Scan()
			nik = input.Text()
			who = fmt.Sprintf("[ %s ]", nik)
			messages <- fmt.Sprintf("%s: has arrived", who)
		default:
			messages <- fmt.Sprintf("%s: %s", who, input.Text())
		}
		//if input.Text() == "EXIT" {
		//	break
		//}
		//messages <- fmt.Sprintf("%s: %s", who, input.Text())
	}
	leaving <- ch
	messages <- fmt.Sprintf("%s: has left", who)
	log.Printf("%s: has left", who)
	err = conn.Close()
	if err != nil {
		return
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		_, err := fmt.Fprintln(conn, msg)
		if err != nil {
			log.Println(err)
		}
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
