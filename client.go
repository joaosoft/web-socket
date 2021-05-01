package web_socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func (client *Client) receive() {
	for {
		message := make([]byte, 4096)
		length, err := client.socket.Read(message)
		if err != nil {
			client.socket.Close()
			break
		}
		if length > 0 {
			fmt.Println("received: " + string(message))
		}
	}
}

func StartClientMode() {
	fmt.Println("starting client...")
	conn, err := net.Dial("tcp", "localhost:12345")
	if err != nil {
		fmt.Println(err)
	}

	client := &Client{socket: conn}

	go client.receive()

		fmt.Printf("message:")
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')
		conn.Write([]byte(strings.TrimRight(message, "\n")))
	}
}
