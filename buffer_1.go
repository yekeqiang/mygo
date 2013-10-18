package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	buffer := bytes.NewBuffer([]byte{0, 1, 2, 3, 4})

	c, _ := buffer.ReadByte()
	buffer.UnreadByte()
	fmt.Println(c)

	for {
		d, err := buffer.ReadByte()
		if err == io.EOF {
			break
		}

		fmt.Println(d)
	}
}
