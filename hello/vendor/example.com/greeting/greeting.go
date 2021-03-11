package greeting

import "fmt"

// Hello module
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome 5!", name)
	return message

}
