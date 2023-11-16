package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	dataStream, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}
	//Using defer to read data stream before the connection closes.
	defer dataStream.Close()

	for {
		connection, err := dataStream.Accept()

		if err != nil {
			panic(err)
		}
		go handle(connection)
	}

}

func handle(connection net.Conn) {

	data, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		panic(err)

	}
	fmt.Printf("data received: %s ", data)
	response := "Message Received!"

	fmt.Fprintf(connection, response+"\n")
	if err != nil {
		panic(err)
	}

	connection.Close()

}
