package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"runtime"
)

func server() {
	listen, _ := net.Listen("tcp", ":8080")

	for {
		conn, _ := listen.Accept()

		go func() {
			defer conn.Close()

			rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

			for {
				s, err := rw.ReadString('\n')

				if err == io.EOF {
					fmt.Println("[server] client close...")
					return
				}

				fmt.Println("[server]", s)
				fmt.Fprintf(rw, "server:%s\n", s)
				rw.Flush()
			}
		}()
	}
}

func client() {
	conn, _ := net.Dial("tcp", "localhost:8080")
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	// fmt.Fprintln(rw,"Hello World!”)
	fmt.Fprintln(rw, "Hello World!")
	rw.Flush()

	s, _ := rw.ReadString('\n')
	fmt.Println("[client]", s)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)
	go server()
	go client()
	<-c
}
