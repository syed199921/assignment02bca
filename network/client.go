package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	// setUpServer()
	setUpClient()

}

func handle(connection net.Conn) {

	data, err := bufio.NewReader(connection).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Printf("data received: %s ", data)

	connection.Close()

}

func setUpServer() {
	dataStream, err := net.Listen("tcp", ":9090")

	if err != nil {
		panic(err)
	}
	//Using defer to read data stream before the connection closes.
	defer dataStream.Close()

	connection, err := dataStream.Accept()

	if err != nil {
		panic(err)
	}
	go handle(connection)
}

func setUpClient() {
	newConnection, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer newConnection.Close()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the message to send: ")
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(newConnection, text+"\n")

	message, err := bufio.NewReader(newConnection).ReadString('\n')
	if err != nil {
		panic(err)

	}
	fmt.Printf("data received: %s ", message)
}
