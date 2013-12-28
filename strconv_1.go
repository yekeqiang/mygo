package main

import (
	"fmt"
	"strconv"
)

func main() {
	bs := make([]byte, 0, 100)
	bs = strconv.AppendInt(bs, 4567, 10)
	bs = strconv.AppendBool(bs, false)
	bs = strconv.AppendQuote(bs, "abcxx")
	bs = strconv.AppendQuoteRune(bs, '我')

	fmt.Println(string(bs))

	fmt.Println(strconv.ParseInt("127", 10, 8))
	fmt.Println(strconv.ParseInt("128", 10, 9))
}
