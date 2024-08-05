package loadtest

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type ResultLoadTest struct {
	StatusCode int
	Duration   time.Duration
}

func StartLoadTest(url string, totalRequests, concurrency int) {
	if concurrency > totalRequests {
		concurrency = totalRequests
		fmt.Printf("Concorrência ajustada para %d (igual ao total de requests)\n", concurrency)
	}

	results := make(chan ResultLoadTest, totalRequests)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	requestsPerWorker := totalRequests / concurrency
	extraRequests := totalRequests % concurrency

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			numRequests := requestsPerWorker
			if workerID < extraRequests {
				numRequests++
			}
			for j := 0; j < numRequests; j++ {
				select {
				case <-ctx.Done():
					return
				default:
					start := time.Now()
					resp, err := http.Get(url)
					duration := time.Since(start)

					if err != nil {
						results <- ResultLoadTest{StatusCode: 0, Duration: duration}
						continue
					}

					results <- ResultLoadTest{StatusCode: resp.StatusCode, Duration: duration}
					resp.Body.Close()
				}
			}
		}(i)
	}

	start := time.Now()
	wg.Wait()
	close(results)

	duration := time.Since(start)
	report(results, duration, totalRequests)
}

func report(results <-chan ResultLoadTest, duration time.Duration, totalRequests int) {
	var totalSuccess, totalFailure int
	statusCodes := make(map[int]int)

	for res := range results {
		if res.StatusCode == 0 {
			totalFailure++
		} else {
			totalSuccess++
			statusCodes[res.StatusCode]++
		}
	}

	fmt.Printf("\nRelatório do Teste de Carga:\n")
	fmt.Printf("Tempo Total: %v\n", duration)
	fmt.Printf("Total de Requests: %d\n", totalRequests)

	// Exibição dos códigos de status HTTP mais importantes
	fmt.Printf("Requests com Status 200 (OK): %d\n", statusCodes[http.StatusOK])
	fmt.Printf("Requests com Status 201 (Created): %d\n", statusCodes[http.StatusCreated])
	fmt.Printf("Requests com Status 202 (Accepted): %d\n", statusCodes[http.StatusAccepted])
	fmt.Printf("Requests com Status 204 (No Content): %d\n", statusCodes[http.StatusNoContent])

	fmt.Printf("Requests com Status 301 (Moved Permanently): %d\n", statusCodes[http.StatusMovedPermanently])
	fmt.Printf("Requests com Status 302 (Found): %d\n", statusCodes[http.StatusFound])

	fmt.Printf("Requests com Status 400 (Bad Request): %d\n", statusCodes[http.StatusBadRequest])
	fmt.Printf("Requests com Status 401 (Unauthorized): %d\n", statusCodes[http.StatusUnauthorized])
	fmt.Printf("Requests com Status 402 (Payment Required): %d\n", statusCodes[http.StatusPaymentRequired])
	fmt.Printf("Requests com Status 403 (Forbidden): %d\n", statusCodes[http.StatusForbidden])
	fmt.Printf("Requests com Status 404 (Not Found): %d\n", statusCodes[http.StatusNotFound])

	fmt.Printf("Requests com Status 500 (Internal Server Error): %d\n", statusCodes[http.StatusInternalServerError])

	// Exibição da distribuição de outros códigos de status HTTP
	fmt.Printf("Distribuição de Outros Códigos de Status HTTP:\n")
	for code, count := range statusCodes {
		if code != http.StatusOK &&
			code != http.StatusCreated &&
			code != http.StatusAccepted &&
			code != http.StatusNoContent &&
			code != http.StatusMovedPermanently &&
			code != http.StatusFound &&
			code != http.StatusBadRequest &&
			code != http.StatusUnauthorized &&
			code != http.StatusPaymentRequired &&
			code != http.StatusForbidden &&
			code != http.StatusNotFound &&
			code != http.StatusInternalServerError {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}

	if totalFailure > 0 {
		fmt.Printf("Requests com Falhas: %d\n", totalFailure)
	}
}
