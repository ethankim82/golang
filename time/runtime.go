package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	for i := 0; i < 100; i++ {
		println("Hello")
	}

	endTime := time.Since(startTime)

	fmt.Printf("runtime %s\n", endTime)
}
