package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	// A slice of web samples
	urls := []string{
		"https://www.easyjet.com",
		"https://www.skyscanner.de",
		"https://www.ryanair.com",
		"https://wizzair.com",
		"https://www.swiss.com",
		"https://testdownfeature.com",
	}
	var wg sync.WaitGroup

	t1 := time.Now()

	for _, u := range urls {
		// Increment the wait group counter
		wg.Add(1)
		go func(url string) {
			// Decrement wait group counter when the go routine completes
			defer wg.Done()
			// Call the actual function
			checkUrl(url)
		}(u)
	}

	// Wait for all check urls to finish
	wg.Wait()
	
	t2 := time.Now()
	fmt.Println("Run time:", t2.Sub(t1))
}

// checks a prints a message if a website is up or down
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down!")
		return
	}

	fmt.Println(url, "is up and running.")
}
