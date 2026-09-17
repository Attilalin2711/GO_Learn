package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

type lesson19Message struct {
	Action string `json:"action"`
	Text   string `json:"text"`
}

// main19 使用 net.Pipe 模擬 TCP 資料流，並以「每行一筆 JSON」切分訊息。
func main19() {
	server, client := net.Pipe()
	defer server.Close()
	defer client.Close()

	go func() {
		encoder := json.NewEncoder(client)
		_ = encoder.Encode(lesson19Message{Action: "login", Text: "Leon"})
		_ = encoder.Encode(lesson19Message{Action: "chat", Text: "Hello"})
		client.Close()
	}()

	// TCP 只提供 byte stream，沒有「一次 Write 就是一次 Read」的保證。
	// Scanner 以換行切割，而 Encoder.Encode 每筆 JSON 都會補上換行。
	scanner := bufio.NewScanner(server)
	for scanner.Scan() {
		var message lesson19Message
		if err := json.Unmarshal(scanner.Bytes(), &message); err != nil {
			fmt.Println("JSON 錯誤：", err)
			continue
		}
		fmt.Printf("action=%s, text=%s\n", message.Action, message.Text)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("讀取資料流失敗：", err)
	}
}
