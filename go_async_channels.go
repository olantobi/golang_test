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

	c := make(chan urlStatus);

	for _, url := range urls {
			go checkUrl(url, c)
	}

	result := make([]urlStatus, len(urls))

	for i, _ := range result {
		result[i] = <- c
		
		if (result[i].status) {
			fmt.Println(result[i].url, "is up")
		} else {
			fmt.Println(result[i].url, "is down!")
		}
	}
	
	t2 := time.Now()
	fmt.Println("Run time:", t2.Sub(t1))
}

// checks a prints a message if a website is up or down
func checkUrl(url string, c chan urlStatus) {
	_, err := http.Get(url)
	if err != nil {
		// Website is down
		c <- urlStatus{url, false}
		return
	}

	// Website is up
	c <- urlStatus{url, true}
}

type urlStatus struct {
	url string
	status bool
}
