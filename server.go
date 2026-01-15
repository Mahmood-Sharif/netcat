package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var (
	clients     = make(map[net.Conn]string)
	history     []string
	mutex       sync.Mutex
	broadcastCh = make(chan string)
	joinCh      = make(chan string)
	leaveCh     = make(chan string)
)

func RunServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}
	defer listener.Close()

	log.Printf("Listening on the port :%s\n", port)

	go handleBroadcast()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection:", err)
			continue
		}

		mutex.Lock()
		if len(clients) >= 10 {
			mutex.Unlock()
			conn.Write([]byte("Server is full. Try again later.\n"))
			conn.Close()
			continue
		}
		mutex.Unlock()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	conn.Write([]byte("Welcome to TCP-Chat!\n"))
	for _, line := range tuxLogo {
		conn.Write([]byte(line + "\n"))
	}
	conn.Write([]byte("[ENTER YOUR NAME]: "))

	scanner := bufio.NewScanner(conn)
	if !scanner.Scan() {
		return
	}
	name := strings.TrimSpace(scanner.Text())
	if name == "" {
		conn.Write([]byte("Name cannot be empty.\n"))
		return
	}

	mutex.Lock()
	clients[conn] = name
	for _, msg := range history {
		conn.Write([]byte(msg + "\n"))
	}
	mutex.Unlock()

	joinCh <- fmt.Sprintf("%s has joined our chat...", name)

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}
		msg := formatMessage(name, text)
		broadcastCh <- msg
		mutex.Lock()
		history = append(history, msg)
		mutex.Unlock()
	}

	leaveCh <- fmt.Sprintf("%s has left our chat...", name)
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
}

func handleBroadcast() {
	for {
		select {
		case msg := <-broadcastCh:
			mutex.Lock()
			for conn := range clients {
				conn.Write([]byte(msg + "\n"))
			}
			mutex.Unlock()
		case joinMsg := <-joinCh:
			mutex.Lock()
			history = append(history, joinMsg)
			for conn := range clients {
				conn.Write([]byte(joinMsg + "\n"))
			}
			mutex.Unlock()
		case leaveMsg := <-leaveCh:
			mutex.Lock()
			history = append(history, leaveMsg)
			for conn := range clients {
				conn.Write([]byte(leaveMsg + "\n"))
			}
			mutex.Unlock()
		}
	}
}

func formatMessage(name, msg string) string {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("[%s][%s]:%s", timestamp, name, msg)
}

var tuxLogo = []string{
	"         _nnnn_",
	"        dGGGGMMb",
	"       @p~qp~~qMb",
	"       M|@||@) M|",
	"       @,----.JM|",
	"      JS^\\__/  qKL",
	"     dZP        qKRb",
	"    dZP          qKKb",
	"   fZP            SMMb",
	"   HZM            MMMM",
	"   FqM            MMMM",
	` __| ".        |\dS"qML`,
	" |    `.       | `' \\Zq",
	"_)      \\.___.,|     .'",
	"\\____   )MMMMMP|   .'",
	"     `-'       `--'",
}
