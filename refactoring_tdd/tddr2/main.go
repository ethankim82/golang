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
	rst := 0
	switch op {

	case "+":
		rst = a + b
	case "-":
		rst = a - b
	case "*":
		rst = a * b
	case "/":
		rst = a / b
	}
	return rst
}
func testCalculate(testcase, op string, a, b, expected int) bool {
	o := calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected:%d output %d \n", testcase, expected, o)
		return false
	}

	return true
}
