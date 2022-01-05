package main

import (
	"fmt"
	"net/http"
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

	t1 := time.Now()

	for _, url := range urls {
			checkUrl(url)
	}
	
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
