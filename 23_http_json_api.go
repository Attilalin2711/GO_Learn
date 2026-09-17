package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type lesson23Book struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type lesson23Error struct {
	Error string `json:"error"`
}

func lesson23WriteJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func lesson23BookHandler(books map[string]lesson23Book) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			lesson23WriteJSON(w, http.StatusMethodNotAllowed, lesson23Error{"只支援 GET"})
			return
		}
		id := r.URL.Query().Get("id")
		if id == "" {
			lesson23WriteJSON(w, http.StatusBadRequest, lesson23Error{"缺少 id"})
			return
		}
		book, ok := books[id]
		if !ok {
			lesson23WriteJSON(w, http.StatusNotFound, lesson23Error{"找不到書籍"})
			return
		}
		lesson23WriteJSON(w, http.StatusOK, book)
	}
}

// main23 將第 16 課 JSON 與 HTTP 結合成查詢 API。
func main23() {
	books := map[string]lesson23Book{"go": {ID: "go", Title: "Go 網路程式設計"}}
	server := httptest.NewServer(lesson23BookHandler(books))
	defer server.Close()

	response, err := server.Client().Get(server.URL + "?id=go")
	if err != nil {
		fmt.Println("請求失敗：", err)
		return
	}
	defer response.Body.Close()
	var book lesson23Book
	if err := json.NewDecoder(response.Body).Decode(&book); err != nil {
		fmt.Println("JSON 解碼失敗：", err)
		return
	}
	fmt.Printf("status=%d, book=%+v\n", response.StatusCode, book)
}
