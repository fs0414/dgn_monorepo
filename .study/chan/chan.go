package main

import "fmt"

func pub(queue chan string) {
	fmt.Println("Hello, World!")
	queue <- "queue"
	close(queue)
}

func sub(queue chan string) {
	queue <- "one"
	queue <- "two"
	queue <- "three"
	fmt.Println(queue)
	<-queue
}

func main() {
	queue := make(chan string)
	go pub(queue)
	go sub(queue)

	fmt.Println("Done!")
}
