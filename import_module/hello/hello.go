package main

import (
	"fmt"
	"math/rand"
	"time"

	"example.com/greeting"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	message := greeting.Hello("Gladys")
	fmt.Println(message, rand.Intn(10))
}
