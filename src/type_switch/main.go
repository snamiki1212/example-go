// https://go.dev/tour/methods/16
package main

import "fmt"

func main() {
	val := randVal()
	switch val.(type) {
	case bool:
		fmt.Println("bool :D", val)
	case string:
		fmt.Println("string :D")
	default:
		fmt.Println("invalid type")
	}
}

func randVal() interface{} {
	return "x"
}
