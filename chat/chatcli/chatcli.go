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
	defer conn.Close()

	go func() {
		io.Copy(os.Stdout, conn)
	}()
		input := bufio.NewScanner(os.Stdin)
		for input.Scan(){
			if input.Text() == "EXIT" {
				fmt.Fprintln(conn, input.Text())
				break
			}
			fmt.Fprintln(conn, input.Text())
		}

	//io.Copy(conn, os.Stdin) // until you send ^Z
	fmt.Printf("%s: exit\n", conn.LocalAddr())
}