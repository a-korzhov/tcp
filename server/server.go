package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
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
		logrus.Println("Message received", message)
		newMessage := strings.ToUpper(message)

		_, err = conn.Write([]byte(newMessage + "\n"))
		if err != nil {
			log.Fatalln(err)
		}
	}

}
