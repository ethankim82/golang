package main

import (
	"fmt"
)

var opMap map[string]func(int, int) int

func initOpMap() {
	opMap = make(map[string]func(int, int) int)
	opMap["+"] = add
	opMap["_"] = sub
	opMap["*"] = mul
	opMap["/"] = div
}
func calculate(op string, a, b int) int {
	if v, ok := opMap[op]; ok {
		return v(a, b)
	}
	return 0
}

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mul(a, b int) int {
	return a * b
}
func div(a, b int) int {
	return a / b
}
func main() {
	test()
}

func test() {
	initOpMap()
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

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected:%d output %d \n", testcase, expected, o)
		return false
	}

	return true
}
