package base

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func OutCall_ct() {
	// 创建socket
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	go broadcast()
	// 监听链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		// 处理连接
		go handleConn(conn)
	}
}

type client chan<- string
var (
	entering = make(chan client)
	leaving = make(chan client)
	message = make(chan string)
)

func broadcast() {
	// 所有的用户
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-message:
			for c := range clients {
				c <- msg
			}
		case c := <-leaving:
			delete(clients, c)
			close(c)
		case c := <-entering:
			clients[c] = true
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	// 绑定用户
	ch := make(chan string)
	// 创建一个goroutine专门给client回数据
	go clientWrite(c, ch)

	who := c.RemoteAddr().String()
	ch <- "you are " + who
	// 上线广播
	entering <- ch
	message <- who + " are arrived"

	// 读取客户端内容
	input := bufio.NewScanner(c)
	for input.Scan() {
		message <- who + ": " + input.Text()
	}

	// 用户下线
	leaving <- ch
	message <- who + " is leave"
}

func clientWrite(c net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(c, msg)
	}
}