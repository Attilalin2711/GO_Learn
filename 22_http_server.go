package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func lesson22Greeting(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "只支援 GET", http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "訪客"
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintln(w, "你好，"+name)
}

// main22 先專注 HTTP 路由、method、query、header 與 status。
func main22() {
	mux := http.NewServeMux()
	mux.HandleFunc("/greet", lesson22Greeting)
	server := httptest.NewServer(mux)
	defer server.Close()

	response, err := server.Client().Get(server.URL + "/greet?name=小明")
	if err != nil {
		fmt.Println("請求失敗：", err)
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("讀取失敗：", err)
		return
	}
	fmt.Println("status：", response.Status)
	fmt.Print("body：", string(body))
}
