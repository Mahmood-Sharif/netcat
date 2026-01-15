package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || len(args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat [server|client] $port")
		return
	}

	mode := args[0]
	port := "8989"
	if len(args) == 2 {
		port = args[1]
	}

	switch mode {
	case "server":
		RunServer(port)
	case "client":
		RunClient(port)
	default:
		fmt.Println("[USAGE]: ./TCPChat [server|client] $port")
	}
}

// RunServer starts the TCP chat server on the specified port.
// RunClient connects to the TCP chat server on the specified port.
