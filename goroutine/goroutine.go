package main

import (
	"fmt"
	"sync"
)

func count(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i, s)
	}
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	/*wg.Add(2)
	go count("fuck")
	wg.Wait()
	deadlock condition
	*/
	wg.Add(1)
	go count("fuck")
	wg.Wait()
}
