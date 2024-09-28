package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// リクエストの中身を表示
	fmt.Println("Request URL:", r.URL.Path)
	fmt.Println("Request Method:", r.Method)
	fmt.Println("Request Header:", r.Header)
	fmt.Println("Request Body:", r.Body)
	w.Write([]byte("Hello, World!"))
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
