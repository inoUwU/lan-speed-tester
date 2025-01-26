package main

import (
	"fmt"
	"net/http"
)

const dataSizeMB = 10 // ダウンロード用のデータサイズ(10MB)

func main() {
	// ダウンロード用のエンドポイント
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		data := make([]byte, dataSizeMB*1024*1024) // ダウンロード用のデータを生成
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
		fmt.Printf("Sent %d MB of data to client\n", dataSizeMB) // 送信したデータのサイズを表示
	})

	// アップロード用のエンドポイント
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		// Post メソッド以外のリクエストを拒否
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		data := make([]byte, r.ContentLength)
		_, err := r.Body.Read(data)
		if err != nil {
			http.Error(w, "Failed to read data", http.StatusInternalServerError)
			return
		}

		fmt.Printf("Received %d bytes from client\n", r.ContentLength) // 受信したデータのサイズを表示
		w.WriteHeader(http.StatusOK)
	})

	// サーバーを起動
	fmt.Println("Starting server on port 8080...")
	fmt.Println("Download: http://localhost:8080/download")
	fmt.Println("Press Ctrl+C to stop server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed", err)
	}
}
