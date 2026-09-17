package main

import (
	"bufio"
	"fmt"
	"net"
)

// main18 在同一個程式內啟動 TCP server 與 client，示範一次完整連線。
func main18() {
	listener, err := net.Listen("tcp", "127.0.0.1:0") // :0 讓系統分配可用 port。
	if err != nil {
		fmt.Println("監聽失敗：", err)
		return
	}
	defer listener.Close()

	serverDone := make(chan error, 1)
	go func() {
		conn, err := listener.Accept()
		if err != nil {
			serverDone <- err
			return
		}
		defer conn.Close()

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			serverDone <- err
			return
		}
		fmt.Fprintln(conn, "伺服器收到："+message)
		serverDone <- nil
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		fmt.Println("連線失敗：", err)
		return
	}
	defer conn.Close()

	fmt.Fprintln(conn, "Hello TCP")
	reply, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("讀取回應失敗：", err)
		return
	}
	fmt.Print(reply)
	if err := <-serverDone; err != nil {
		fmt.Println("伺服器錯誤：", err)
	}
}
