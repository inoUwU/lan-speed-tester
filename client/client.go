package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"sync"
	"time"
)

const (
	downloadURL     = "http://localhost:8080/download"
	dataSizeMB      = 10 // データサイズ（MB）
	numMeasurements = 5  // 測定回数
	threads         = 4  // 並列ダウンロードのスレッド数
)

// 並列ダウンロード速度測定
func parallelDownload(url string, threads int) float64 {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			defer resp.Body.Close()
			io.Copy(io.Discard, resp.Body) // データを捨てる
		}()
	}

	wg.Wait()
	duration := time.Since(start).Seconds()
	totalData := float64(dataSizeMB*threads) * 8 // データ量（ビット）
	return totalData / (duration * 1024 * 1024)  // Mbpsで返す
}

// 測定結果を分析
func analyzeSpeeds(speeds []float64) (float64, float64) {
	sort.Float64s(speeds)
	return calculateAverage(speeds), calculateMedian(speeds)
}

func calculateAverage(speeds []float64) float64 {
	var total float64
	for _, speed := range speeds {
		total += speed
	}
	return total / float64(len(speeds))
}

func calculateMedian(speeds []float64) float64 {
	mid := len(speeds) / 2
	if len(speeds)%2 == 0 {
		return (speeds[mid-1] + speeds[mid]) / 2
	}
	return speeds[mid]
}

func main() {
	fmt.Println("Measuring download speed...")
	speeds := make([]float64, numMeasurements)

	// 複数回測定
	for i := 0; i < numMeasurements; i++ {
		speeds[i] = parallelDownload(downloadURL, threads)
		fmt.Printf("Measurement %d: %.2f Mbps\n", i+1, speeds[i])
	}

	// 測定結果の分析
	avg, median := analyzeSpeeds(speeds)
	fmt.Printf("\nAverage Speed: %.2f Mbps\n", avg)
	fmt.Printf("Median Speed: %.2f Mbps\n", median)
}
