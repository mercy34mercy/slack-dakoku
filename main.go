package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストの中身を表示
	fmt.Println("Request URL:", r.URL.Path)
	fmt.Println("Request Method:", r.Method)
	fmt.Println("Request Header:", r.Header)

	// リクエストボディを読み込む
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// リクエストボディの中身を表示
	fmt.Println("Request Body:", string(body))

	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
