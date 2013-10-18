package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, _ := os.Open("chan_1.go")
	defer f.Close()

	r := bufio.NewReaderSize(f, 4096)
	for {
		line, isPrefix, err := r.ReadLine()

		if isPrefix {
			// fmt.Println(string(line))
			fmt.Println(line)
		} else if len(line) > 0 {
			// fmt.Println(string(line))
			fmt.Println(line)
		}

		if err == io.EOF {
			break
		}

	}
}
