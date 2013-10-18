package main

import (
	"fmt"
)

type Data struct {
	args []int
	ch   chan string
}

func server() chan *Data {
	reqs := make(chan *Data) //服务器Request Channel

	go func() { //使用独立的goroutine 循环处理所有的客户端请求
		for d := range reqs { //循环从Request Channel获取客户端
			go serverProcess(d) //将请求的客户端数据提交给另外的
		}
	}()

	return reqs
}

func serverProcess(data *Data) {
	x := 0
	for i := range data.args {
		x += i
	}

	s := fmt.Sprintf("server:%d", x)

	data.ch <- s
}

func main() {
	/* 启动服务器 */

	serverReqs := server()

	/* 客户端向服务器发送请求数据，并等待返回结果 */

	data := &Data{[]int{1, 2, 3, 4}, make(chan string)}
	serverReqs <- data     //发送请求到服务器
	fmt.Println(<-data.ch) //获取服务器返回结果

	/*  关闭服务器 */
	close(serverReqs)
}
