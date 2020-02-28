package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		log.Print("Send message: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		log.Print("Message from server: ", message)
	}
}
