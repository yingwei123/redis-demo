package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/exp/rand"
)

func main() {
	baseURL := "http://localhost:3000" // Replace with your actual base URL
	iter := 100
	fmt.Printf("Test Speed with getting 1 product %d times\n", iter)
	CompareGetProductSpeed(baseURL, iter)

	fmt.Printf("Test Speed with getting all products %d times\n", iter)
	CompareGetAllProductSpeed(baseURL, iter)

	fmt.Printf("Test Speed with updating 1 product %d times\n", iter)
	CompareUpdateProductSpeed(baseURL, iter)
}

func CompareGetAllProductSpeed(baseURL string, numRequests int) {
	// Test without Redis
	start := time.Now()
	for i := 0; i < numRequests; i++ {
		resp, err := http.Get(baseURL + "/products")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		resp.Body.Close()
	}

	withoutRedisTime := time.Since(start)

	// Test with Redis
	start = time.Now()
	for i := 0; i < numRequests; i++ {
		resp, err := http.Get(baseURL + "/products/redis")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		resp.Body.Close()
	}
	withRedisTime := time.Since(start)

	// Print results
	fmt.Printf("Time taken for %d requests without Redis: %v\n", numRequests, withoutRedisTime)
	fmt.Printf("Time taken for %d requests with Redis: %v\n", numRequests, withRedisTime)
	fmt.Printf("Speed improvement: %.2f%%\n", (float64(withoutRedisTime-withRedisTime)/float64(withoutRedisTime))*100)
}

func CompareUpdateProductSpeed(baseURL string, numRequests int) {
	productID := "1"

	// Helper function to generate random price and update product
	updateProductPrice := func(url string) error {
		// Generate a random price between 1 and 200
		randomPrice := rand.Intn(200) + 1

		// Create the product update payload
		productUpdate := map[string]interface{}{
			"price": randomPrice,
		}

		// Marshal the product update to JSON
		body, err := json.Marshal(productUpdate)
		if err != nil {
			return fmt.Errorf("error marshalling request body: %v", err)
		}

		// Make a PUT request with the JSON body
		req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		// Send the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("error making request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("non-200 response: %d", resp.StatusCode)
		}

		return nil
	}

	// Test without Redis
	start := time.Now()
	for i := 0; i < numRequests; i++ {
		err := updateProductPrice(baseURL + "/product/" + productID)
		if err != nil {
			fmt.Printf("Error making request without Redis: %v\n", err)
			return
		}
	}
	withoutRedisTime := time.Since(start)

	// Test with Redis
	start = time.Now()
	for i := 0; i < numRequests; i++ {
		err := updateProductPrice(baseURL + "/product/" + productID + "/redis")
		if err != nil {
			fmt.Printf("Error making request with Redis: %v\n", err)
			return
		}
	}
	withRedisTime := time.Since(start)

	// Print results
	fmt.Printf("Time taken for %d requests without Redis: %v\n", numRequests, withoutRedisTime)
	fmt.Printf("Time taken for %d requests with Redis: %v\n", numRequests, withRedisTime)
	fmt.Printf("Speed improvement: %.2f%%\n", (float64(withoutRedisTime-withRedisTime)/float64(withoutRedisTime))*100)
}

func CompareGetProductSpeed(baseURL string, numRequests int) {
	productID := "1"

	// Test without Redis
	start := time.Now()
	for i := 0; i < numRequests; i++ {
		resp, err := http.Get(baseURL + "/product/" + productID)
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		resp.Body.Close()
	}

	withoutRedisTime := time.Since(start)

	// Test with Redis
	start = time.Now()
	for i := 0; i < numRequests; i++ {
		resp, err := http.Get(baseURL + "/product/" + productID + "/redis")
		if err != nil {
			fmt.Printf("Error making request: %v\n", err)
			return
		}
		resp.Body.Close()
	}
	withRedisTime := time.Since(start)

	// Print results
	fmt.Printf("Time taken for %d requests without Redis: %v\n", numRequests, withoutRedisTime)
	fmt.Printf("Time taken for %d requests with Redis: %v\n", numRequests, withRedisTime)
	fmt.Printf("Speed improvement: %.2f%%\n", (float64(withoutRedisTime-withRedisTime)/float64(withoutRedisTime))*100)
}
