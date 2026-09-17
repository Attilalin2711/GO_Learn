package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

type lesson24Result struct {
	Message string `json:"message"`
}

// main24 示範建立 request、設定 header、timeout、檢查狀態與解碼 body。
func main24() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header.Get("X-API-Key") != "demo-key" {
			w.WriteHeader(http.StatusUnauthorized)
			_ = json.NewEncoder(w).Encode(lesson23Error{Error: "無效 API key"})
			return
		}
		_ = json.NewEncoder(w).Encode(lesson24Result{Message: "驗證成功"})
	}))
	defer server.Close()

	request, err := http.NewRequest(http.MethodGet, server.URL, nil)
	if err != nil {
		fmt.Println("建立 request 失敗：", err)
		return
	}
	request.Header.Set("X-API-Key", "demo-key")
	client := &http.Client{Timeout: 2 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("傳送失敗：", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP 錯誤：", response.Status)
		return
	}
	var result lesson24Result
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		fmt.Println("解碼失敗：", err)
		return
	}
	fmt.Printf("status=%s, message=%s\n", response.Status, result.Message)
}
