package main

import (
	"fmt"
)

func main() {
	//UTF-8 编码格式存储的字符串
	s := "a:中国人"
	fmt.Println("utf-8 string:", s, len(s))

	/* ------- string to bytes -------------*/
	//转换成bytes，以便于修改
	bs := []byte(s)
	bs[0] = 'A'

	for i := 0; i < len(bs); i++ {
		fmt.Printf("%02x", bs[i])
	}

	//将bytes转换成string
	fmt.Println(string(bs))

	/* --------- string to unicode -------------- */
	//转换成unicode
	u := []rune(s)
	fmt.Println("unicode len=", len(u))

	//显示unicode number
	for i := 0; i < len(u); i++ {
		fmt.Printf("%04x", u[i])
	}

	//按照unicode 修改
	u[4] = '龙'
	//将unicode 装换成string
	fmt.Println(string(u))

}
