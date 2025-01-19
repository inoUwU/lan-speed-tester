package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	serverURL := "http://localhost:8080"

	// ダウンロード速度測定
	fmt.Println("Measuring download speed...")
	start := time.Now()
	resp, err := http.Get(serverURL + "/download")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer resp.Body.Close()
	byteDownloaded, _ := io.Copy(io.Discard, resp.Body)
	duration := time.Since(start).Seconds()
	downloadSpeed := float64(byteDownloaded*8) / (1024 * 1024) / duration
	fmt.Printf("Download speed: %.2f Mbps\n", downloadSpeed)
}
