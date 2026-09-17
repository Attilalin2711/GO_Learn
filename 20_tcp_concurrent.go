package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

func lesson20Handle(conn net.Conn) {
	defer conn.Close()
	name, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return
	}
	fmt.Fprint(conn, "歡迎，"+name)
}

// main20 為每一個連線啟動 goroutine，使 server 能同時服務多個 client。
func main20() {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		fmt.Println("監聽失敗：", err)
		return
	}
	defer listener.Close()

	const clientCount = 3
	var handlers sync.WaitGroup
	serverDone := make(chan struct{})
	go func() {
		defer close(serverDone)
		for i := 0; i < clientCount; i++ {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			handlers.Add(1)
			go func() {
				defer handlers.Done()
				lesson20Handle(conn)
			}()
		}
		handlers.Wait()
	}()

	var clients sync.WaitGroup
	for i := 1; i <= clientCount; i++ {
		clients.Add(1)
		go func(id int) {
			defer clients.Done()
			conn, err := net.Dial("tcp", listener.Addr().String())
			if err != nil {
				fmt.Println("client 連線失敗：", err)
				return
			}
			defer conn.Close()
			fmt.Fprintf(conn, "client-%d\n", id)
			reply, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print(reply)
		}(i)
	}
	clients.Wait()
	<-serverDone
}
