package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func writeLog(line string) {
	file, err := os.OpenFile(
		"honeypot.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(line + "\n")
}

func logAndPrint(message string) {
	fmt.Println(message)
	writeLog(message)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	remote := conn.RemoteAddr().String()
	logAndPrint("Handling connection from" + remote)

	conn.Write([]byte("SSH-2.0-OpenSSH_8.2p1 Ubuntu-4ubuntu0.5\r\n"))

	reader := bufio.NewReader(conn)

	conn.Write([]byte("login: "))
	username, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	conn.Write([]byte("password: "))
	password, err := reader.ReadString('\n')
	if err != nil {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	logLine := fmt.Sprintf(
		"[%s] %s | login=%s password=%s",
		timestamp,
		remote,
		strings.TrimSpace(username),
		strings.TrimSpace(password),
	)
	logAndPrint(logLine)

	conn.Write([]byte("Authentication failed.\r\n"))

}

func main() {

	listener, err := net.Listen("tcp", ":2222")
	if err != nil {
		fmt.Println("Failed to open port:", err)
		return
	}

	defer listener.Close()

	fmt.Println("Honeypot Listening on port 2222")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}
}
