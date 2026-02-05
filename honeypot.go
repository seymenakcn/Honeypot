package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("Handling connection from", conn.RemoteAddr())

	conn.Write([]byte("SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.5\r\n"))

	reader := bufio.NewReader(conn)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println("Recieved", input)
}

func main() {
	fmt.Println("Honeypot Starting...")

	listener, err := net.Listen("tcp", ":2222")
	if err != nil {
		fmt.Println("Failed to open port:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Listening on port 2222")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Println("New Connection from", conn.RemoteAddr())
		go handleConnection(conn)
	}
}
