package web_socket

import (
	"fmt"
	"net"
)

func StartServerMode() {
	fmt.Println("starting server...")
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println(err)
	}

	manager := ClientManager{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go manager.start()

	for {
		fmt.Println("waiting for clients communication...")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		client := &Client{socket: conn, data: make(chan []byte)}
		manager.register <- client

		go manager.receive(client)
		go manager.send(client)
	}
}
