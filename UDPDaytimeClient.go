package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if lent(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	updAddr, err := net.ResolveUDPAddr("up4", service)

	checkError(err)

	conn, err := net.DialUDP("udp", nil, updAddr)
	checkError(err)

	_, err = conn.Write([]byte("anything"))
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(string(buf[0:n]))

	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Fatal error", err.Error())
		os.Exit(1)
	}
}
