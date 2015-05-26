package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var arrayClients = make(map[int]net.Conn)

func handleConn(conn net.Conn, c chan string, outUser chan int, position int) {
	defer conn.Close()
	//Get Username
	conn.Write([]byte("Insira seu nome:"))
	username, _, _ := bufio.NewReader(conn).ReadLine()
	nick := string(username)
	//Get Message
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			x := nick + " saiu da sala\n"
			c <- x
			outUser <- position
			break
		}
		x := nick + ": " + string(message)
		c <- x
	}
}

func sendMsg(message string) {
	//Echo message to All Clients
	for _, value := range arrayClients {
		value.Write([]byte(message))
	}
}

func listenChannel(c chan string, inUser chan net.Conn, outUser chan int) {
	for {
		select {
		case msg1 := <-c:
			sendMsg(msg1)
		case client := <-inUser:
			arrayClients[len(arrayClients)] = client
			go handleConn(client, c, outUser, len(arrayClients)-1)
		case deleteClient := <-outUser:
			delete(arrayClients, deleteClient)
		}
	}
}

func main() {
	c := make(chan string)
	inUser := make(chan net.Conn)
	outUser := make(chan int)
	listener, err := net.Listen("tcp", ":13000")
	if err != nil {
		fmt.Println("Listen: ", err)
		os.Exit(1)
	}

	go listenChannel(c, inUser, outUser)

	for {
		fmt.Println("Esperando conexão...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept: ", err)
			os.Exit(1)
		}
		inUser <- conn
	}
}
