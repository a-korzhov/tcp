package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	fmt.Println("Starting server...")

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Message received", message)
		newMessage := strings.ToUpper(message)

		_, err = conn.Write([]byte(newMessage + "\n"))
		if err != nil {
			log.Fatalln(err)
		}
	}

}
