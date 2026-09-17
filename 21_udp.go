package main

import (
	"fmt"
	"net"
	"time"
)

// main21 示範 UDP datagram。UDP 不需 Accept，也不保證送達、順序或重傳。
func main21() {
	server, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	if err != nil {
		fmt.Println("UDP 監聽失敗：", err)
		return
	}
	defer server.Close()
	_ = server.SetReadDeadline(time.Now().Add(2 * time.Second))

	done := make(chan error, 1)
	go func() {
		buffer := make([]byte, 1024)
		n, clientAddress, err := server.ReadFromUDP(buffer)
		if err != nil {
			done <- err
			return
		}
		_, err = server.WriteToUDP([]byte("ACK: "+string(buffer[:n])), clientAddress)
		done <- err
	}()

	client, err := net.DialUDP("udp4", nil, server.LocalAddr().(*net.UDPAddr))
	if err != nil {
		fmt.Println("UDP 連線設定失敗：", err)
		return
	}
	defer client.Close()
	_ = client.SetDeadline(time.Now().Add(2 * time.Second))

	_, _ = client.Write([]byte("temperature=25.6"))
	buffer := make([]byte, 1024)
	n, err := client.Read(buffer)
	if err != nil {
		fmt.Println("UDP 接收失敗：", err)
		return
	}
	fmt.Println(string(buffer[:n]))
	if err := <-done; err != nil {
		fmt.Println("UDP server 錯誤：", err)
	}
}
