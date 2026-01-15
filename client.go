package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func RunClient(port string) {
	conn, err := net.Dial("tcp", "localhost:"+port)
	if err != nil {
		fmt.Printf("❌ Connection error: %v\n", err)
		return
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		text, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("❌ Error reading input:", err)
			break
		}
		text = strings.TrimSpace(text)
		if text == "" {
			continue
		}
		_, err = fmt.Fprintln(conn, text)
		if err != nil {
			fmt.Println("❌ Error sending message:", err)
			break
		}
	}
}
