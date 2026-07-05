package main

import (
	"errors"
	"fmt"
	"os"
)

/*
error 是 Go 內建的錯誤型別
*/
func divid(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("zero divided error")
	}

	return a / b, nil
}

/*
defer 是在目前 function 裡註冊一個延後執行的動作，等到這個 function 準備結束時，才依照 後註冊先執行 的順序執行
defer 採LIFO
readFile()是一個典型讀檔典型案例
*/
func deferdemo() {
	defer fmt.Println("close database")
	defer fmt.Println("close file")
	defer fmt.Println("unlock mutex")

	fmt.Println("doing work:=========")
}

func readFile() error {
	file, err := os.Open("test.txt")
	if err != nil {
		return err
	}

	defer file.Close()

	// 假設這裡發生錯誤，提前 return
	return errors.New("read failed")

	// 即使提前 return，file.Close() 還是會執行
}

/*
panic = 主動丟出嚴重錯誤，開始中止目前流程
recover = 在 defer 裡接住 panic，阻止它繼續往外傳

經典用例:
Server 處理每個 request 時，用 recover 防止單一 request 的 panic 把整個服務打掛
*/
func handleRequest() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("request failed, recovered:", r)
		}
	}()

	fmt.Println("start request")

	// 假設這裡發生不可預期錯誤
	panic("database connection nil")

	fmt.Println("finish request") // 不會執行
}

func main08() {

	result, err := divid(10, 0)
	//Go 最常見的錯誤處理模式
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	fmt.Println("=======defer demo========")
	deferdemo()

	fmt.Println("=======panic+recovery========")

	handleRequest()
	fmt.Println("server still running")
}
