package main

import (
	"bytes"
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

	// アップロード速度測定
	fmt.Print("Measuring upload speed...")
	data := bytes.Repeat([]byte("A"), 1024*1024)
	start = time.Now()
	resp, err = http.Post(serverURL+"/upload", "application/octet-stream", bytes.NewReader(data))
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer resp.Body.Close()
	duration = time.Since(start).Seconds()
	uploadSpeed := float64(len(data)*8) / (1024 * 1024) / duration
	fmt.Printf("Upload speed: %.2f Mbps\n", uploadSpeed)
}
