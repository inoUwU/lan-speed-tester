package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ダウンロード用のエンドポイント
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		data := make([]byte, 10*1024*1024) // 10MBのダミーデータを作成
		w.Write(data)
	})

	// アップロード用のエンドポイント
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %d bytes\n", r.ContentLength) // 受信したデータのサイズを表示
		w.WriteHeader(http.StatusOK)
	})

	fmt.Println("Starting server on port 8080...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed", err)
	}
}
