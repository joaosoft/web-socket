package main

import (
	"flag"
	"strings"
	"web-socket"
)

func main() {
	flagMode := flag.String("mode", "server", "start in client or server mode")
	flag.Parse()
	if strings.ToLower(*flagMode) == "server" {
		web_socket.StartServerMode()
	} else {
		web_socket.StartClientMode()
	}
}