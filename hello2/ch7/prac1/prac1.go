package main

import "fmt"

func main() {
	var res = Multiple(10, 20)

	fmt.Println(res)
}

func Multiple(a, b int) int {
	return a * b
}
