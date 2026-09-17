package main

import (
	"fmt"
	"net"
)

// main17 示範網路程式最常見的 address、host、port 與 IP 解析。
func main() {
	//格視為host:port,其中host為loopback,也就是本機ip
	address := "127.0.0.1:8080"
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		fmt.Println("拆解位址失敗：", err)
		return
	}
	fmt.Printf("完整位址=%s，host=%s，port=%s\n", address, host, port)

	// JoinHostPort 比字串相加安全，尤其能正確處理 IPv6 位址。
	fmt.Println("重新組合：", net.JoinHostPort(host, port))

	// ParseIP 只解析 IP 字串，不進行 DNS 查詢。
	ip := net.ParseIP("127.0.0.1")
	if ip == nil {
		fmt.Println("不是合法 IP")
		return
	}
	fmt.Println("是否為 loopback：", ip.IsLoopback())

	// ResolveTCPAddr 將文字位址轉成 TCPAddr；tcp4 表示限定 IPv4。
	tcpAddress, err := net.ResolveTCPAddr("tcp4", address)
	if err != nil {
		fmt.Println("解析 TCP 位址失敗：", err)
		return
	}
	fmt.Printf("TCP 位址：IP=%s，Port=%d\n", tcpAddress.IP, tcpAddress.Port)
}
