package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	url         string
	requests    int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&requests, "requests", 0, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 0, "Número de chamadas simultâneas")
}

func main() {
	flag.Parse()

	if url == "" || requests <= 0 || concurrency <= 0 {
		fmt.Println("Parâmetros inválidos. Use --url, --requests e --concurrency.")
		return
	}

	var wg sync.WaitGroup
	requestsChan := make(chan int, requests)
	resultsChan := make(chan int, requests)

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(&wg, requestsChan, resultsChan)
	}

	for i := 0; i < requests; i++ {
		requestsChan <- i
	}
	close(requestsChan)

	wg.Wait()
	close(resultsChan)

	totalTime := time.Since(startTime)
	totalRequests := 0
	successfulRequests := 0
	statusCodes := make(map[int]int)

	for status := range resultsChan {
		totalRequests++
		if status == http.StatusOK {
			successfulRequests++
		}
		statusCodes[status]++
	}

	fmt.Printf("Tempo total gasto: %v\n", totalTime)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successfulRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for code, count := range statusCodes {
		fmt.Printf("Status %d: %d\n", code, count)
	}
}

func worker(wg *sync.WaitGroup, requestsChan <-chan int, resultsChan chan<- int) {
	defer wg.Done()
	for range requestsChan {
		resp, err := http.Get(url)
		if err != nil {
			resultsChan <- 0
			continue
		}
		resultsChan <- resp.StatusCode
		resp.Body.Close()
	}
}
