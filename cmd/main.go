package main

import "fmt"

func main() {
	fmt.Println(Sum(255, 255))
}

// Sum ...
func Sum(a, b byte) byte {
	return a + b
}
