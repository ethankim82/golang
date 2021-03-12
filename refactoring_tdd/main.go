package main

import "fmt"

func main() {
	test()
}
func test() {
	o := calculate("+", 2, 3)
	if o != 5 {
		fmt.Printf("Test Failed expected %d outpupt %d", 5, o)
		return
	}
	o = calculate("+", 1, -3)
	if o != -2 {
		fmt.Printf("Test Failed expected %d outpupt %d", 5, o)
		return
	}
	o = calculate("-", 1, 3)
	if o != -2 {
		fmt.Printf("Test Failed expected %d outpupt %d", 5, o)
		return
	}
	fmt.Printf("Success!")
}
func calculate(op string, a, b int) int {
	if op == "+" {
		return a + b
	}
	if op == "-" {
		return a- b
	}
	return 0

}
