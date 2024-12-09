package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 暫定的なヘルスチェックエンドポイント
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
