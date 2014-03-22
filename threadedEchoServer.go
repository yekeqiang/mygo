package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("ip4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)

	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		_, err := conn.Write(buf[0:n])
		if err != nil {
			return
		}

		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}