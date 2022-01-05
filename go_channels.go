package main

import "fmt"

func sum(a int, b int, c chan int) {
	c <- a + b
}

func main() {
	c := make(chan int)
	go sum(20, 22, c)
	x := <- c

	fmt.Println(x)
}