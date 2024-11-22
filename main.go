package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 100, "Número total de requests")
	concurrency := flag.Int("concurrency", 10, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("A URL deve ser fornecida.")
		return
	}

	var wg sync.WaitGroup
	statusCodes := make(map[int]int)
	start := time.Now()

	sem := make(chan struct{}, *concurrency)

	for i := 0; i < *totalRequests; i++ {
		sem <- struct{}{}
		wg.Add(1)

		go func() {
			defer wg.Done()
			resp, err := http.Get(*url)
			if err != nil {
				fmt.Println("Erro ao fazer a requisição:", err)
				<-sem
				return
			}
			statusCodes[resp.StatusCode]++
			defer resp.Body.Close()
			_, _ = ioutil.ReadAll(resp.Body)
			<-sem
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)

	fmt.Printf("Tempo total: %s\n", elapsed)
	fmt.Printf("Total de requests: %d\n", *totalRequests)
	fmt.Printf("Requests com status 200: %d\n", statusCodes[200])
	for code, count := range statusCodes {
		if code != 200 {
			fmt.Printf("Requests com status %d: %d\n", code, count)
		}
	}
}
