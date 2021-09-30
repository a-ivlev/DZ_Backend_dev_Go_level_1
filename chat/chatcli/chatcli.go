package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			log.Println(err)
		}
	}(conn)

	go func() {
		_, err = io.Copy(os.Stdout, conn)
		if err != nil {
			log.Println(err)
			return
		}
	}()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "EXIT" {
			_, err = fmt.Fprintln(conn, input.Text())
			if err != nil {
				log.Println(err)
				return
			}
			break
		}
		_, err = fmt.Fprintln(conn, input.Text())
		if err != nil {
			log.Println(err)
			return
		}
	}

	fmt.Printf("%s: exit\n", conn.LocalAddr())
}