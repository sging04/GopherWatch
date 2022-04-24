package main

import (
	"fmt"
	"sync"
	"time"
)

func countApples() {
	var i int = 0
	for {
		fmt.Printf("Apple number %v\n", i)
		time.Sleep(400 * time.Millisecond)
		i++
	}
}

func countOranges() {
	var i int = 0
	for {
		fmt.Printf("Orange number %v\n", i)
		time.Sleep(500 * time.Millisecond)
		i++
	}
}

func main() {
	var wg sync.WaitGroup /* Insantiating WaitGroup */

	wg.Add(2) /* Telling it to expect 2 go-routines within its scope */

	go countApples()
	go countOranges()

	wg.Wait() /* Waiting for all routines to end */
}
