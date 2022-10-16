package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func UpdateMessage(s string) {
	defer wg.Done()
	msg = s
}

func main() {
	msg = "Hello World!"
	wg.Add(1)

	go UpdateMessage("Hello, Universe")
	wg.Wait()
	fmt.Println(msg)
	wg.Add(1)
	go UpdateMessage("Hello, Cosmos")

	wg.Wait()

	fmt.Println(msg)
}
