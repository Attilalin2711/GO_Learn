package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestLessonExamples 依序執行所有範例，確認本機 client/server 能正常結束。
func TestLessonExamples(t *testing.T) {
	lessons := []struct {
		name string
		run  func()
	}{
		{"17_network_basic", main17},
		{"18_tcp", main18},
		{"19_tcp_protocol", main19},
		{"20_tcp_concurrent", main20},
		{"21_udp", main21},
		{"22_http_server", main22},
		{"23_http_json_api", main23},
		{"24_http_client", main24},
	}
	for _, lesson := range lessons {
		t.Run(lesson.name, func(t *testing.T) {
			lesson.run()
		})
	}
}

// TestLesson23BookHandler 不開真實監聽埠，直接驗證 handler 的回應。
func TestLesson23BookHandler(t *testing.T) {
	books := map[string]lesson23Book{"go": {ID: "go", Title: "Go 網路程式設計"}}
	tests := []struct {
		name       string
		method     string
		target     string
		wantStatus int
	}{
		{name: "查詢成功", method: http.MethodGet, target: "/?id=go", wantStatus: http.StatusOK},
		{name: "缺少 id", method: http.MethodGet, target: "/", wantStatus: http.StatusBadRequest},
		{name: "找不到", method: http.MethodGet, target: "/?id=missing", wantStatus: http.StatusNotFound},
		{name: "method 錯誤", method: http.MethodPost, target: "/?id=go", wantStatus: http.StatusMethodNotAllowed},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(test.method, test.target, nil)
			recorder := httptest.NewRecorder()
			lesson23BookHandler(books).ServeHTTP(recorder, request)
			if recorder.Code != test.wantStatus {
				t.Fatalf("status=%d, want=%d", recorder.Code, test.wantStatus)
			}
			var body map[string]any
			if err := json.NewDecoder(recorder.Body).Decode(&body); err != nil {
				t.Fatalf("回應不是合法 JSON：%v", err)
			}
		})
	}
}
