package main

import (
	"bufio" //implements buffered I/O
	"fmt"
	"log"
	"net" //package for networking I/O
)

func main() {
	startTCPSever()
}

func startTCPSever() {
	//create a listener
	myListener, err := net.Listen("tcp", ":5331")
	if err != nil {
		fmt.Println(err)
		log.Fatalln(err)
		return
	}
	defer myListener.Close()

	for {
		//accept connection
		conn, err := myListener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//handle connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	//create a new scanner
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Received: ", string(data))

	retMsg := string(data) + "\n"
	conn.Write([]byte(retMsg))
}
