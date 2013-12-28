package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "Hello World!")

	buffer := bytes.NewBuffer(nil)
	fmt.Fprintf(buffer, "%x\n", h.Sum(nil))

	fmt.Println(buffer.String())
}
